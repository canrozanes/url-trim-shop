package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect initiates a mongo connection and returns the client
func Connect(mongoURI string) (*mongo.Client, context.CancelFunc, context.Context) {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("could not connect to mongo %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("error connecting to mongo client %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	return client, cancel, ctx
}
