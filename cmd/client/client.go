package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ajvideira/fullcycle-grpc/pb"
	"google.golang.org/grpc"
)


func main() {

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection);

	addUser(client)

}

func addUser(client pb.UserServiceClient) {
	req := &pb.User{
		Name: "Jonathan",
		Email: "jonathan.videira@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req);
	if err != nil {
		log.Fatalf("Could not call AddUser: %v", err)
	}

	fmt.Println(res)
}