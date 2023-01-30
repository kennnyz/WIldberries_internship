package main

import (
	"context"
	_ "example.com/go-usermgmt-grpc/usermgmt"
	pb "example.com/go-usermgmt-grpc/usermgmt"
	"fmt"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Didn't connect %v", err)
	}
	defer conn.Close()
	c := pb.NewUsermgmtClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newUsers = make(map[string]int32)
	newUsers["Alice"] = 43
	newUsers["Alex"] = 23
	newUsers["Sam"] = 19
	newUsers["John"] = 24

	for name, age := range newUsers {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("Could not create a new user: %v", err)
		}
		fmt.Printf(`User details: 
NAME: %s
AGE: %d
ID: %d
`, r.GetName(), r.GetAge(), r.GetId())
	}

}
