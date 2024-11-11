package main

import (
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	usecaseBooks "bookservice/business/books"
	usecaseBooksStocks "bookservice/business/booksstocks"
	handlerBooks "bookservice/controller/books"
	handlerBooksStocks "bookservice/controller/booksstocks"
	repoBooks "bookservice/databases/books"
	repoBooksStocks "bookservice/databases/booksstocks"
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
	booksStocksUseCase := initBooksStocksServer(db)
	handlerBooks.NewServer(grpcServer, bookUseCase)
	handlerBooksStocks.NewServer(grpcServer, booksStocksUseCase)

	log.Fatal(grpcServer.Serve(lis))
}

func initBookServer(db *gorm.DB) usecaseBooks.BookUseCaseInterface {
	bookRepo := repoBooks.NewBookRepository(db)
	return usecaseBooks.NewUseCase(bookRepo)
}

func initBooksStocksServer(db *gorm.DB) usecaseBooksStocks.BooksStocksUseCaseInterface {
	booksStocksRepo := repoBooksStocks.NewBooksStocksRepository(db)
	return usecaseBooksStocks.NewUseCase(booksStocksRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&repoBooks.Book{}, &repoBooksStocks.BooksStocks{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
