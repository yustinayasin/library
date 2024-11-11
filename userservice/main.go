package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	middleware "userservice/app/middlewares"
	usecase "userservice/business"
	handler "userservice/controller"
	repo "userservice/databases"
	"userservice/helpers"
)

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	secretJWT := os.Getenv("JWT_SECRET")
	expiresDuration := os.Getenv("JWT_EXPIRED")
	expiresDurationInt, _ := strconv.Atoi(expiresDuration)

	db, err := helpers.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	migrations(db)

	jwtConf := middleware.ConfigJWT{
		SecretJWT:       secretJWT,
		ExpiresDuration: expiresDurationInt,
	}

	// add a listener address
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	// start the grpc server
	grpcServer := grpc.NewServer()

	userUseCase := initUserServer(db, jwtConf)
	handler.NewServer(grpcServer, userUseCase)

	// start serving to the address
	log.Fatal(grpcServer.Serve(lis))
}

func initUserServer(db *gorm.DB, jwtConf usecase.GeneratorToken) usecase.UserUseCaseInterface {
	userRepo := repo.NewUserRepository(db)
	return usecase.NewUseCase(userRepo, jwtConf)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&repo.User{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
