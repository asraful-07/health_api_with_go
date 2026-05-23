package handler

import (
	"context"
	"encoding/json"
	"go-project/models"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func SetDB(conn *pgx.Conn) {
	DB = conn
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Health checker web site!"))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id"
	err = DB.QueryRow(context.Background(), query, newUser.Name, newUser.Email, newUser.Age).Scan(&newUser.Id)

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, email, age FROM users"
	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
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

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	query := "SELECT id, name, email, age FROM users WHERE id = $1"
	var user models.User
	err = DB.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Name, &user.Email, &user.Age)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4"
	_, err = DB.Exec(context.Background(), query, updatedUser.Name, updatedUser.Email, updatedUser.Age, id)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	updatedUser.Id = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM users WHERE id = $1"
	_, err = DB.Exec(context.Background(), query, id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


