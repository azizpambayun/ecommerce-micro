package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/azizpambayun/ecommerce-micro/models"
)

var users = []models.User{
	{ID:1, Username: "ilham", Email: "ilham@windah.com", Password: "hashedpassword", Created: time.Now()},
}

// GetUserHandler handles fetching user information
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// simulate fetching a user
	if len(users) > 0 {
		json.NewEncoder(w).Encode(users)
	} else {
		http.Error(w, "No users found", http.StatusNotFound)
	}
}

// AddUserHandler handles adding a new user
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid reques body", http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	newUser.Created = time.Now()
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}