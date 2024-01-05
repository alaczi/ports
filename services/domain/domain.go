package main

import (
	"context"
	"fmt"
	pb "github.com/alaczi/ports/ports"
	repo "github.com/alaczi/ports/repository"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
)

type Config struct {
	Port int `default:"50051"`
}

type PortDomainService struct {
	pb.UnimplementedPortDomainServiceServer
	repository repo.PortRepository
}

func newServer() *PortDomainService {
	repository := newInMemoryPortRepository()
	s := &PortDomainService{
		repository: repository,
	}
	return s
}

func (s *PortDomainService) GetPort(_ context.Context, id *pb.PortId) (*pb.Port, error) {
	port, err := s.repository.GetPort(id.Id)
	if err != nil {
		return &pb.Port{}, err
	}
	if port == nil {
		return &pb.Port{}, status.Error(codes.NotFound, "Port not found")
	}
	return pb.ToProtoPort(port), nil
}

func (s *PortDomainService) UpsertPort(_ context.Context, port *pb.Port) (*pb.Empty, error) {
	return &pb.Empty{}, s.repository.UpsertPort(pb.ToPort(port))
}

func (s *PortDomainService) UpsertPorts(stream pb.PortDomainService_UpsertPortsServer) error {
	var receivedTotal uint64
	for {
		port, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.PortSummary{
				ReceivedTotal: receivedTotal,
			})
		}
		if err != nil {
			return err
		}
		err = s.repository.UpsertPort(pb.ToPort(port))
		if err != nil {
			return err
		}
		receivedTotal++
	}
}

func main() {
	var c Config
	err := envconfig.Process("domain", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Service started on port: %v", c.Port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPortDomainServiceServer(grpcServer, newServer())
	grpcServer.Serve(listener)
}