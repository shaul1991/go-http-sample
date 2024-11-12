package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go-http/internal/route"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", env))
	if err != nil {
		log.Fatalf("Error loading .env file for %s environment", env)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set in the environment")
	}

	router := route.SetupRoutes() // Use the setup function to initialize routes
	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
