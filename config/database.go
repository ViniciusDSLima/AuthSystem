package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var db *mongo.Database

func ConnectDatabase() error {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}

	clientOpts := options.Client().
		ApplyURI(mongoURL).
		SetRetryWrites(true).
		SetRetryReads(true).
		SetCompressors([]string{"zstd"}).
		SetServerSelectionTimeout(time.Second * 30)

	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer pingCancel()

	err = client.Ping(pingCtx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB!")

	db = client.Database(os.Getenv("MONGO_DB_NAME"))
	if db == nil {
		return fmt.Errorf("database not found")
	}

	return nil
}

func GetCollection(name string) *mongo.Collection {
	return db.Collection(name)
}
