package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

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

func connectDB() {
 var err error
 connStr := "postgres://postgres:rahat12@localhost:5432/go_crud_db"

 db, err = pgx.Connect(context.Background(), connStr)
 if err != nil {
	 panic(err)
 }
}

func main() {
	connectDB()
	defer db.Close(context.Background())

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

query := "INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id"
err = db.QueryRow(context.Background(), query, newUser.Name, newUser.Email, newUser.Age).Scan(&newUser.Id)

if err != nil {
	http.Error(w, "Error creating user", http.StatusInternalServerError)
	return
}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func getsUsersHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, email, age FROM users"
	rows, err := db.Query(context.Background(), query)
if err != nil {
	http.Error(w, "Error fetching users", http.StatusInternalServerError)
	return
}

defer rows.Close()

var users []User

for rows.Next() {
	var user User
	err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age)
	if err != nil {
		http.Error(w, "Error scanning user", http.StatusInternalServerError)
		return
	}	
	users = append(users, user)
    }

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

query := "UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4"
_, err = db.Exec(context.Background(), query, updatedUser.Name, updatedUser.Email, updatedUser.Age, id)
if err != nil {
	http.Error(w, "Error updating user", http.StatusInternalServerError)
	return
}

updatedUser.Id = id
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(updatedUser)
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
		// users = append(users[:i], users[i+1:]...)
		users = slices.Delete(users, i, i+1)

		w.WriteHeader(http.StatusNoContent)
		return
	}	
		
}
http.Error(w, "User not found", http.StatusNotFound)
}

