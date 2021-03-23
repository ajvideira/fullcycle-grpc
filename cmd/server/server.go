package main

import (
	"log"
	"net"

	"github.com/ajvideira/fullcycle-grpc/pb"
	"github.com/ajvideira/fullcycle-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {

	lis, err := net.Listen("tcp", "localhost:50051");
	if err != nil {
		log.Fatalf("Could not listen: %v", err);
	}

	grpcServer := grpc.NewServer();

	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis);
	if err != nil {
		log.Fatalf("Could not serve: %v", err);
	}

}