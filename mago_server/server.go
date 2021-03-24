package main

import (
	"context"
	"flag"
	"fmt"
	mago "github.com/alirezastack/mago/magopb"
	"github.com/nyaruka/phonenumbers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

var port = flag.Int("port", 50051, "gRPC Listening Port")

type server struct {
	mago.UnimplementedMagoServiceServer
}

func (s *server) CreateUser(ctx context.Context, in *mago.CreateUserRequest) (*mago.CreateUserResponse, error) {
	log.Printf("Received user phone: %v", in.GetPhone())
	// We just check client cancellation before expensive calls
	//if ctx.Err() == context.Canceled {
	//	fmt.Println("The client canceled the request!")
	//	return nil, status.Error(codes.Canceled, "The client canceled the request!")
	//}

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
	flag.Parse()
	fmt.Printf("Starting up Mago server on 0.0.0.0:%v...\n", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// initialize encryption
	certFile := "helpers/ssl/server.crt"
	keyFile := "helpers/ssl/server.pem"
	creds, certErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if certErr != nil {
		log.Fatalf("Failed loading certificates: %v", certErr)
	}

	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)
	mago.RegisterMagoServiceServer(s, &server{})

	fmt.Println("Mago server is up & running ;)")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server on listener: %v", err)
	}
}
