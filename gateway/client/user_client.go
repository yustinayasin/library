package users

import (
	"log"
	"shared/proto"

	"google.golang.org/grpc"
)

func NewUserServiceClient(address string) proto.UserServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to User Service: %v", err)
	}
	return proto.NewUserServiceClient(conn)
}
