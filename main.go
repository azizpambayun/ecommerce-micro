package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/azizpambayun/ecommerce-micro/router"
)

func main() {
	r := router.IntializeRouter()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }

    fmt.Printf("Starting User Service on port %s...\n", port)
    if err := http.ListenAndServe(":"+port, r); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }

}