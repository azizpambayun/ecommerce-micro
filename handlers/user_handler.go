package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/azizpambayun/ecommerce-micro/models"
	"github.com/segmentio/kafka-go"
)

var users = []models.User{
	{ID:1, Username: "ilham", Email: "ilham@windah.com", Password: "hashedpassword", Created: time.Now()},
}

// Kafka writer to send messages
var kafkaWriter = &kafka.Writer{
	Addr: kafka.TCP("localhost:9902"),
	Topic: "user-events",
	Balancer: &kafka.LeastBytes{},
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


	// Publish the event to Kafka
	msg := fmt.Sprintf("New user added: %s", newUser.Username)
	err := kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Key: []byte(fmt.Sprintf("user-%d", newUser.ID)),
		Value: []byte(msg),
	}) 
	if err != nil {
		log.Printf("could not write message to kafka: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}