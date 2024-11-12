package main

import (
	"fmt"
	"gateway/routes"
	"log"
	"os"
	middleware "shared/app/middlewares"
	protoBook "shared/proto/books"
	protoUser "shared/proto/users"
	"strconv"

	"github.com/joho/godotenv"
	"go-micro.dev/v5"
	"google.golang.org/grpc"
)

type Gateway struct {
	UserClient         protoUser.UserServiceClient
	BookClient         protoBook.BookServiceClient
	BookStockClient    protoBook.BooksStocksServiceClient
	BorrowRecordClient protoUser.BorrowRecordsServiceClient
}

func NewGateway(userServiceAddr, bookServiceAddr string) (*Gateway, error) {
	userConn, err := grpc.Dial(userServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %v", err)
	}
	userClient := protoUser.NewUserServiceClient(userConn)
	borrowRecordClient := protoUser.NewBorrowRecordsServiceClient(userConn)

	bookConn, err := grpc.Dial(bookServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to book service: %v", err)
	}
	bookClient := protoBook.NewBookServiceClient(bookConn)
	bookStockClient := protoBook.NewBooksStocksServiceClient(bookConn)

	return &Gateway{
		UserClient:         userClient,
		BookClient:         bookClient,
		BookStockClient:    bookStockClient,
		BorrowRecordClient: borrowRecordClient,
	}, nil
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

	gateway, err := NewGateway("localhost:5001", "localhost:5002")

	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}

	routes.RouteRegister(jwtConf, gateway.UserClient, gateway.BookClient, gateway.BookStockClient, gateway.BorrowRecordClient)
}
