package main

import (
	"fmt"
	"net/http"

	"github.com/kelvinator05/clean-architecture-go/internal/infrastructure/db"
	"github.com/kelvinator05/clean-architecture-go/internal/infrastructure/server"
	"github.com/kelvinator05/clean-architecture-go/internal/usecase"
)

func main() {
	userRepo := db.NewInMemoryUserRepo()
	userUseCase := usecase.UserUseCase{Repo: userRepo}
	httpServer := server.NewHTTPServer(userUseCase)

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", httpServer)
	mux.Handle("/users", httpServer)
	mux.Handle("/users/", httpServer)

	// Run the server
	fmt.Println("Server running on port: 8080")
	http.ListenAndServe(":8080", mux)
}
