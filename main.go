package main

import (
	"log"
	"net/http"
	"os"

	"go-http/route"

	"go-http/middleware"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatalf("Error loading .env file for %s environment", env)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set in the environment")
	}

	r := route.SetupRoutes()
	r.Use(middleware.ErrorHandler) // Use the error handling middleware

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
