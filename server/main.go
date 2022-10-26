package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/alisavch/grpc-service/pkg/api"
)

type server struct {
	api.UnimplementedHasherServer
}

func newServer() *server {
	s := &server{}
	return s
}

func (s *server) Convert(ctx context.Context, request *api.InputNote) (*api.OutputNote, error) {
	encoding := base64.StdEncoding.EncodeToString([]byte(request.Message))
	output := &api.OutputNote{
		Message: encoding,
	}
	return output, nil
}

func main() {
	var opts []grpc.ServerOption

	lis, err := net.Listen("tcp", fmt.Sprintf(":8081"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server is running...")

	grpcServer := grpc.NewServer(opts...)
	api.RegisterHasherServer(grpcServer, newServer())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("Server is disabled")
}
