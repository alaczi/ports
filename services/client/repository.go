package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/alaczi/ports/ports"
	repo "github.com/alaczi/ports/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"
)

type GRPCPortRepository struct {
	client pb.PortDomainServiceClient
}

func newGRPCPortRepository(portServiceAddr *string, dataFile *string) *GRPCPortRepository {
	conn, err := grpc.Dial(*portServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the GRPC endpoint: %v", err)
	}
	client := pb.NewPortDomainServiceClient(conn)
	repo := &GRPCPortRepository{
		client: client,
	}
	log.Print("Connected to the grpc endpoint")
	err = repo.initializeData(dataFile)
	if err != nil {
		log.Fatalf("Data upload failed: %v", err)
	}
	log.Print("Initial data upload finished")
	return repo
}

func (s *GRPCPortRepository) UpsertPort(port *repo.Port) error {
	return errors.New("Not supported")
}

func (s *GRPCPortRepository) GetPort(id string) (*repo.Port, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := s.client.GetPort(ctx, &pb.PortId{Id: id})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.NotFound {
				return nil, nil
			}
		}
		log.Printf("client.GetPort failed: %v", err)
		return nil, err
	}
	return pb.ToPort(result), nil
}

func (s *GRPCPortRepository) initializeData(dataFile *string) error {
	f, err := os.Open(*dataFile)
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("Error while closing the json file: %v", err)
		}
	}()
	dec := json.NewDecoder(f)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := s.client.UpsertPorts(ctx)
	if err != nil {
		log.Printf("client.UpsertPorts failed: %v", err)
		return err
	}
	defer func() {
		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("client.UpsertPorts failed: %v", err)
		}
		log.Printf("Uploaded ports summary: %v", reply)
	}()
	// read open bracket from the json file
	_, err = dec.Token()
	if err != nil {
		log.Print(err)
		return err
	}
	// read each top level property
	for dec.More() {
		// read the next token which is supposed to be the name of the property which holds the port data
		token, err := dec.Token()
		if err != nil {
			log.Print(err)
			return err
		}
		var id string
		switch token.(type) {
		case string:
			id = fmt.Sprintf("%v", token)
			break
		default:
			return nil
		}

		var port repo.Port
		// decode the value of the property
		err = dec.Decode(&port)
		port.Id = id
		if err != nil {
			log.Print(err)
			return err
		}
		if err := stream.Send(pb.ToProtoPort(&port)); err != nil {
			log.Printf("client.UpsertPorts: stream.Send(%v) failed: %v", port.Id, err)
			return err
		}
	}

	// read the closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
