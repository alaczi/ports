package api

import (
	"context"
	"fmt"
	pb "github.com/alaczi/ports/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"port_domain_service/services"
)

type PortServer struct {
	pb.UnimplementedPortDomainServiceServer
	portService services.PortServiceInterface
	port        int
	server      *grpc.Server
}

func NewPortServer(config *services.Config, portService services.PortServiceInterface) *PortServer {
	s := &PortServer{
		portService: portService,
		port:        config.Port,
	}
	return s
}

func (s *PortServer) Serve() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Service started on port: %v", s.port)
	var opts []grpc.ServerOption
	s.server = grpc.NewServer(opts...)
	pb.RegisterPortDomainServiceServer(s.server, s)
	s.server.Serve(listener)
}

func (s *PortServer) Shutdown() {
	s.server.GracefulStop()
}

func (s *PortServer) GetPort(ctx context.Context, id *pb.PortId) (*pb.Port, error) {
	port, err := s.portService.GetPort(ctx, id.Id)
	if err != nil {
		return &pb.Port{}, err
	}
	if port == nil {
		return &pb.Port{}, status.Error(codes.NotFound, "Port not found")
	}
	return pb.ToProtoPort(port), nil
}

func (s *PortServer) UpsertPort(ctx context.Context, port *pb.Port) (*pb.Empty, error) {
	return &pb.Empty{}, s.portService.UpsertPort(ctx, pb.ToPort(port))
}

func (s *PortServer) UpsertPorts(stream pb.PortDomainService_UpsertPortsServer) error {
	var receivedTotal uint64
	ctx := stream.Context()
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
		err = s.portService.UpsertPort(ctx, pb.ToPort(port))
		if err != nil {
			return err
		}
		receivedTotal++
	}
}