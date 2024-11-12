package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	minPoolSize     uint64 = 5
	maxPoolSize     uint64 = 100
	connectTimeout         = 10 * time.Second
	maxConnIdleTime       = 30 * time.Second
)

// Client is a MongoDB client instance
var Client *mongo.Client

// Connect establishes a connection to MongoDB with connection pool
func Connect(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	// Configure connection pool options
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize).
		SetMaxConnIdleTime(maxConnIdleTime)

	// Create client with connection pool
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	Client = client
	log.Printf("Connected to MongoDB! Pool size: %d-%d", minPoolSize, maxPoolSize)
	return nil
}

// GetCollection returns a collection from the database
func GetCollection(dbName, collectionName string) *mongo.Collection {
	return Client.Database(dbName).Collection(collectionName)
}

// Disconnect closes the MongoDB connection pool
func Disconnect() error {
	if Client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	err := Client.Disconnect(ctx)
	if err != nil {
		return err
	}

	log.Println("Disconnected from MongoDB")
	return nil
} 