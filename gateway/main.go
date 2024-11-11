package main

import (
	"gateway/routes"
	"log"
	"os"
	middleware "shared/app/middlewares"
	"shared/proto"
	"strconv"

	"github.com/joho/godotenv"
	"go-micro.dev/v5"
	"google.golang.org/grpc"
)

type Gateway struct {
	UserClient proto.UserServiceClient
}

func NewGateway(userServiceAddr string) (*Gateway, error) {
	conn, err := grpc.Dial(userServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	userClient := proto.NewUserServiceClient(conn)
	return &Gateway{UserClient: userClient}, nil
}

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	secretJWT := os.Getenv("JWT_SECRET")
	expiresDuration := os.Getenv("JWT_EXPIRED")
	expiresDurationInt, _ := strconv.Atoi(expiresDuration)

	jwtConf := middleware.ConfigJWT{
		SecretJWT:       secretJWT,
		ExpiresDuration: expiresDurationInt,
	}

	service := micro.NewService(
		micro.Name("Gateway"),
		micro.Address(":8080"),
	)

	service.Init()

	gateway, err := NewGateway("localhost:5001")

	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}

	routes.RouteRegister(jwtConf, gateway.UserClient)
}
