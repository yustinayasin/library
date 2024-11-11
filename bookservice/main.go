package main

import (
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	usecase "bookservice/business"
	handler "bookservice/controller"
	repo "bookservice/databases"
	"bookservice/helpers"
)

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := helpers.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	migrations(db)

	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	grpcServer := grpc.NewServer()

	bookUseCase := initBookServer(db)
	handler.NewServer(grpcServer, bookUseCase)

	log.Fatal(grpcServer.Serve(lis))
}

func initBookServer(db *gorm.DB) usecase.BookUseCaseInterface {
	bookRepo := repo.NewBookRepository(db)
	return usecase.NewUseCase(bookRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&repo.Book{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
