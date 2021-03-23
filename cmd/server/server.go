package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)


func main() {

	lis, err := net.Listen("tcp", "localhost:50051");
	if err != nil {
		log.Fatalf("Could not listen: %v", err);
	}

	grpcServer := grpc.NewServer();

	err = grpcServer.Serve(lis);
	if err != nil {
		log.Fatalf("Could not serve: %v", err);
	}

}