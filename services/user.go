package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ajvideira/fullcycle-grpc/pb"
)


type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.GetName())

	return &pb.User{
		Id: "123",
		Name: req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (userService *UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Preparing Insert",
		User: nil,
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User: nil,
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Insert Complete",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})

	return nil
}

func (userService *UserService) AddUsers(stream pb.UserService_AddUsersServer) error {

	users := []*pb.User{}

	index := 0;

	for {
		index = index + 1;

		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("End of request stream")
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving request stream: %v", err)
		}
		fmt.Println("Receiving user ", req.GetName())
		users = append(users, &pb.User{
			Id: fmt.Sprint(index),
			Name: req.GetName(),
			Email: req.GetEmail(),
		})
	}

}

func (userService *UserService) AddUsersBidirectional(stream pb.UserService_AddUsersBidirectionalServer) (error) {
	index := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("End of receive streaming")
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving data: %v", err)
		}

		time.Sleep(time.Second * 3)

		fmt.Printf("Init insert process for %s\n", req.GetName())
		stream.Send(&pb.UserResultStream{
			Status: fmt.Sprintf("Init insert process for %s", req.GetName()),
			User: nil,
		})

		time.Sleep(time.Second * 3)

		fmt.Printf("Inserting %s\n", req.GetName())
		stream.Send(&pb.UserResultStream{
			Status: fmt.Sprintf("Inserting %s", req.GetName()),
			User: nil,
		})
		
		time.Sleep(time.Second * 3)

		index = index + 1
		fmt.Printf("%s complete insert\n", req.GetName())
		stream.Send(&pb.UserResultStream{
			Status: fmt.Sprintf("%s complete insert", req.GetName()),
			User: &pb.User{
				Id: fmt.Sprint(index),
				Name: req.GetName(),
				Email: req.GetEmail(),
			},
		})
	}
}