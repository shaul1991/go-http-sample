package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go-http/internal/database/mongodb"
	"go-http/internal/database/mysql"
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

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI must be set in the environment")
	}

	err = mongodb.Connect(mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongodb.Disconnect()

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	err = mysql.Connect(mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer mysql.Disconnect()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set in the environment")
	}

	router := route.SetupRoutes()
	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
