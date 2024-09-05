package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() *mongo.Client {
	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		log.Fatalf("MONGODB_URL environment variable not set")
	}
	log.Printf("Connecting to MongoDB using URI: %s", mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB!")
	MongoClient = client
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database("fastify").Collection(collectionName)
}
