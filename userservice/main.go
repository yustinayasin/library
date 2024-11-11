package main

import (
	"fmt"
	"log"
	"net/http"

	"go-micro.dev/v5"
)

func main() {
	// Create a new web service
	service := micro.NewService(
		micro.Name("UserService"),
		micro.Address(":8080"),
	)

	// Initialize the service
	service.Init()

	// Set up a route and handler
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
