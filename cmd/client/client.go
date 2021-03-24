package main

import (
	"context"
	"fmt"
	"io"
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

	client := pb.NewUserServiceClient(connection)

	addUser(client)
	addUserVerbose(client)
}

func addUser(client pb.UserServiceClient) {
	req := &pb.User{
		Name: "Jonathan",
		Email: "jonathan.videira@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not call AddUser: %v", err)
	}

	fmt.Println(res)
}

func addUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Name: "Jonathan",
		Email: "jonathan.videira@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not call AddUser: %v", err)
	}

	for {
		stream, err := responseStream.Recv();
		if err == io.EOF {
			fmt.Println("End of stream.")
			break
		} 
		if err != nil {
			log.Fatalf("Error receiving the stream: %v", err)
		}
		fmt.Println(stream.GetStatus())
	}
}