package main

import (
	"log"
	"net"

	"github.com/rmcaxoul/grpc_server/calculator"
	"google.golang.org/grpc"
)

// Starts the server on port 9000 and initializes gRPC
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Could not listen to port 9000: %v", err)
	}

	s := calculator.Server{}

	grpcServer := grpc.NewServer()

	calculator.RegisterCalculatorServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server port 9000: %s", err)
	}
}
