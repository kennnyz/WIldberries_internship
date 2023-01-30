package main

import (
	"context"
	pb "example.com/go-usermgmt-grpc/usermgmt"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUsermgmtServer
}

//CreateNewUser(context.Context, *NewUser) (*User, error)

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	fmt.Printf("Received: %v \n", in.GetName())
	var userID int32 = int32(rand.Intn(10000))
	return &pb.User{
		Id:   userID,
		Name: in.GetName(),
		Age:  in.GetAge(),
	}, nil
}

func main() {
	// listen on TCP port 50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a gRPC server object
	s := grpc.NewServer()
	pb.RegisterUsermgmtServer(s, &UserManagementServer{})

	fmt.Printf("Server is listening on port %s \n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
