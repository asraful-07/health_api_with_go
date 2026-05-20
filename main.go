package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
    mux.HandleFunc("GET /user/{id}", getUserHandler)
    mux.HandleFunc("PUT /user/{id}", updateUserHandler)
    mux.HandleFunc("DELETE /user/{id}", deleteUserHandler)

	fmt.Println("Server running on :8080")
	 err := http.ListenAndServe(":8080", mux)
	 
	 if err != nil {
		 fmt.Println("Error starting server:", err)
	 }
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Health checker web site!"))	
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
idParam := r.PathValue("id")

id, err := strconv.Atoi(idParam)
if err != nil {
	http.Error(w, "Invalid user ID", http.StatusBadRequest)
	return
}
for _, user := range users {
	if user.Id == id {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
		return
	}
}
http.Error(w, "User not found", http.StatusNotFound)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
idParam := r.PathValue("id")
id, err := strconv.Atoi(idParam)
if err != nil {
	http.Error(w, "Invalid user ID", http.StatusBadRequest)
	return
}

var updatedUser User
err = json.NewDecoder(r.Body).Decode(&updatedUser)
if err != nil {
	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	return
}

for i, user := range users {
	if user.Id == id {
		updatedUser.Id = id
		users[i] = updatedUser	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedUser)
		return
	}
}
http.Error(w, "User not found", http.StatusNotFound)
}


func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
idParam := r.PathValue("id")
id, err := strconv.Atoi(idParam)
if err != nil {
	http.Error(w, "Invalid user ID", http.StatusBadRequest)
	return
}
for i, user := range users {
	if user.Id == id {
		users = append(users[:i], users[i+1:]...)

		w.WriteHeader(http.StatusNoContent)
		return
	}	
		
}
http.Error(w, "User not found", http.StatusNotFound)
}