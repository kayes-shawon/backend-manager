package database

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var db *mongo.Database

type MongoDB struct {
	db *mongo.Database
}

func Setup() {
	_, err := mongoConnect()
	if err != nil {
		log.Fatal(err)
	}
}

func mongoConnect() (*MongoDB, error) {
	connectionString := viper.GetString("DB_CONNECTION_URL")
	databaseName := viper.GetString("DATABASE_NAME")

	// Set up a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to the MongoDB database
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the MongoDB server to check if it's reachable.
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB server: %v", err)
	}

	db = client.Database(databaseName)
	return &MongoDB{db}, nil
}
