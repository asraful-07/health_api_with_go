package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id int  `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

var users = []User{
	{
		Id: 1, 
		Name: "Alice", 
		Email: "alice@example.com", 
		Age: 30,
	},
	{
		Id: 2,
		Name: "Bob",
		Email: "bob@example.com",
		Age: 23,
	},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /users", getsUsersHandler)
	fmt.Println("Server running on :8080")
	 err := http.ListenAndServe(":8080", mux)
	 
	 if err != nil {
		 fmt.Println("Error starting server:", err)
	 }
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))	
	
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser.Id = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func getsUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}