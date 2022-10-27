package main

import (
	"fmt"
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
)

type lovooServer struct {}

func (s *lovooServer) Add (ctx context.Context, req *OpRequest) (OpResponse, error){
	return nil, nil
}

func main(){
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}