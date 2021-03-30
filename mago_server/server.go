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
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var port = flag.Int("port", 50051, "gRPC Listening Port")

type RpcServer struct {
	mago.UnimplementedMagoServiceServer
	address    string
	grpcServer *grpc.Server
}

func NewServer(addr string) *RpcServer {
	// initialize encryption
	certFile := "helpers/ssl/server.crt"
	keyFile := "helpers/ssl/server.pem"
	creds, certErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if certErr != nil {
		log.Fatalf("Failed loading certificates: %v", certErr)
	}

	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)

	rs := &RpcServer{
		address:    addr,
		grpcServer: s,
	}
	return rs
}

func (s *RpcServer) CreateUser(ctx context.Context, in *mago.CreateUserRequest) (*mago.CreateUserResponse, error) {
	log.Printf("Received user phone: %v\n", in.GetPhone())
	// We just check client cancellation before expensive calls
	//if ctx.Err() == context.Canceled {
	//	log.Println("The client canceled the request!")
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

func (s *RpcServer) Run() {
	listener, err := net.Listen("tcp", s.address)
	defer listener.Close()

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("rpc listening on: %v rpc_server\n", s.address)

	mago.RegisterMagoServiceServer(s.grpcServer, s)
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server on listener: %v", err)
	}
}

// It stops the server from accepting new connections and RPCs and
// blocks until all the pending RPCs are finished.
func (s *RpcServer) Stop() {
	log.Printf("Kill signal received, stopping Mago server...\n")
	s.grpcServer.GracefulStop()
	log.Printf("Mago server is stopped!\n")
}

// go run main.go --port 8085
func main() {
	// on go code crash, we receive file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()
	log.Printf("Starting up Mago server on 0.0.0.0:%v...\n", *port)

	server := NewServer(fmt.Sprintf("0.0.0.0:%d", *port))
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		server.Run()
		wg.Done()
	}()

	// Signal handling and graceful shutdown of gRPC server
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	server.Stop()

	wg.Wait()
}
