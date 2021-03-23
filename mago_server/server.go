package main

import (
	"context"
	"fmt"
	mago "github.com/alirezastack/mago/magopb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	mago.UnimplementedMagoServiceServer
}

// CreateUser implements mago.CreateUser
func (s *server) CreateUser(ctx context.Context, in *mago.CreateUserRequest) (*mago.CreateUserResponse, error) {
	log.Printf("Received user phone: %v", in.GetPhone())
	return &mago.CreateUserResponse{UserId: "FAKE-ID"}, nil
}

func main() {
	fmt.Println("Starting up Mago server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	mago.RegisterMagoServiceServer(s, &server{})

	fmt.Println("Mago server is up & running ;)")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server on listener: %v", err)
	}
}
