package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//addUser(client)
	//addUserVerbose(client)
	addUsers(client)
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

func addUsers(client pb.UserServiceClient) {
	requestStream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error sending the stream: %v", err)
	}
	err = requestStream.Send(&pb.User{
		Name: "Jonathan",
		Email: "jonathan.videira@gmail.com",
	})
	if err != nil {
		log.Fatalf("Error sending the user Jonathan: %v", err)
	}

	time.Sleep(time.Second * 3)

	err = requestStream.Send(&pb.User{
		Name: "Manuela",
		Email: "castilhos.manuela@gmail.com",
	})
	if err != nil {
		log.Fatalf("Error sending the user Manuela: %v", err)
	}

	users, err := requestStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing de send stream: %v", err)
	}
	fmt.Println(users)
}