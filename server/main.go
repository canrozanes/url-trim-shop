package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"url-trimmer/config"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const port = ":5000"

type URLHashPair struct {
	Hash string `json: hash`
	URL  string `json: "url"`
}

// InMemoryURLStore is temporary implementation
type MongoURLStore struct {
	hashes *mongo.Collection
}

// HashURL hashes a given url
func (m *MongoURLStore) HashURL(url string) string {
	return "abc"
}

func (m *MongoURLStore) AddHashToStore(hash string, url string) {
	res, err := m.hashes.InsertOne(context.Background(), bson.M{"url": url, "hash": hash})
	if err != nil {
		log.Fatalf("could not add hash: %s and url: %s, %v", hash, url, err)
	}
	// TODO use id to create hash
	_ = res
	return
}

// GetHashFromURL is temporary implementation
func (m *MongoURLStore) GetHashFromURL(url string) string {
	urlHashPair := URLHashPair{}
	filter := bson.D{{"url", url}}
	err := m.hashes.FindOne(context.Background(), filter).Decode(&urlHashPair)
	if err == mongo.ErrNoDocuments {
		newHash := m.HashURL(url)
		m.AddHashToStore(newHash, url)
		return newHash
	} else if err != nil {
		log.Fatalf("could not get hash for url: %s, %v", url, err)
	}
	return urlHashPair.Hash
}

func (m *MongoURLStore) GetURLFromHash(hash string) string {
	urlHashPair := URLHashPair{}
	filter := bson.D{{"hash", hash}}
	err := m.hashes.FindOne(context.Background(), filter).Decode(&urlHashPair)
	if err == mongo.ErrNoDocuments {
		return ""
	} else if err != nil {
		log.Fatalf("could not get url for hash: %s, %v", hash, err)
		return ""
	}
	return urlHashPair.URL
}

// NewMongoServer initiates a HashingServer connected to MongoDB
func NewMongoServer(client *mongo.Client) *HashingServer {
	return &HashingServer{&MongoURLStore{
		hashes: client.Database("url-trimmer").Collection("hashes"),
	}}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := config.Connect()
	server := NewMongoServer(client)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port %s %v", port, err)
	}

	fmt.Printf("Listening on port %s", port)
}
