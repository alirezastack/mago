package main

import (
	"context"
	"fmt"
	mago "github.com/alirezastack/mago/magopb"
	"github.com/nyaruka/phonenumbers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type server struct {
	mago.UnimplementedMagoServiceServer
}

// CreateUser implements mago.CreateUser
func (s *server) CreateUser(ctx context.Context, in *mago.CreateUserRequest) (*mago.CreateUserResponse, error) {
	log.Printf("Received user phone: %v", in.GetPhone())
	num, err := phonenumbers.Parse(in.GetPhone(), "IR")
	formattedNumber := phonenumbers.Format(num, phonenumbers.E164)
	log.Printf("International Phone Number: %v", formattedNumber)
	if err != nil {
		log.Printf("Failed to format phone number %v: %v", in.GetPhone(), err)
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid phone number given: %v", in.GetPhone()),
		)
	}

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
