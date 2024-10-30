package router

import (
	"github.com/azizpambayun/ecommerce-micro/handlers"
	"github.com/gorilla/mux"
)

func IntializeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/add", handlers.AddUserHandler).Methods("POST")
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	return r
}