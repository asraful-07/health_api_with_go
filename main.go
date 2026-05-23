package main

import (
	"context"
	"fmt"
	"net/http"

	"go-project/db"
	"go-project/handler"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	db.ConnectDB()
	handler.SetDB(db.DB)
	defer db.DB.Close(context.Background())

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /", handler.RootHandler)
	mux.HandleFunc("POST /users", handler.CreateUserHandler)
	mux.HandleFunc("GET /users", handler.GetUsersHandler)
	mux.HandleFunc("GET /user/{id}", handler.GetUserHandler)
	mux.HandleFunc("PUT /user/{id}", handler.UpdateUserHandler)
	mux.HandleFunc("DELETE /user/{id}", handler.DeleteUserHandler)

	fmt.Println("Server running on :8080")
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

