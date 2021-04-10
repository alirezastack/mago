package main

import (
	"context"
	"fmt"
	mago "github.com/alirezastack/mago/magopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func RpcClient(address string, service string) *grpc.ClientConn {
	fmt.Printf("%v client is about to call a remote method...", service)
	certFile := "helpers/ssl/ca.crt" // Certificate Authority Trust Certificate
	creds, SSLErr := credentials.NewClientTLSFromFile(certFile, "")
	if SSLErr != nil {
		log.Fatalf("Error in TLS CA Cert initiation: %v", SSLErr)
	}

	cc, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("grpc Dial error: %v", err)
	}
	return cc
}

func main() {
	cc := RpcClient("localhost:50051", "mago")
	defer cc.Close()
	c := mago.NewMagoServiceClient(cc)
	fmt.Println("Starting to do a Unary RPC...")
	in := &mago.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "09111111111",
	}
	// Timeout after 200ms
	ctx, cancelFunc := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelFunc()

	res, err := c.CreateUser(ctx, in)
	if err != nil {
		respError, ok := status.FromError(err)
		if ok {
			//	actual error from gRPC (user error)
			fmt.Printf("Error message from Mago Server: %v\n", respError.Message())
			fmt.Printf("Error code from Mago Server: %v\n", respError.Code())
			if respError.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline was exceeded!")
			}
		} else {
			log.Fatalf("User not created: %v", err)
		}
		return
	}

	log.Printf("User created: %v", res.UserId)
}
