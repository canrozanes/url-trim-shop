package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"url-trimmer/config"
	"url-trimmer/utils"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const port = ":5000"

// URLHashPair stores a URL and its Hash
type URLHashPair struct {
	Hash string
	URL  string
}

// MongoURLStore stores a collection of URLHashPairs
type MongoURLStore struct {
	hashes *mongo.Collection
}

// HashURL hashes a given url
func (m *MongoURLStore) HashURL(url string) string {
	countOfRecords, err := m.hashes.CountDocuments(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		return ""
	}

	newHash := utils.ToBase62(uint64(countOfRecords) + 1)

	return newHash
}

// AddHashToStore adds hash to store
func (m *MongoURLStore) AddHashToStore(hash string, url string) {
	_, err := m.hashes.InsertOne(context.Background(), bson.M{"url": url, "hash": hash})
	if err != nil {
		log.Fatalf("could not add hash: %s and url: %s, %v", hash, url, err)
	}
	return
}

// GetHashFromURL is temporary implementation
func (m *MongoURLStore) GetHashFromURL(url string) string {
	urlHashPair := URLHashPair{}
	filter := bson.D{{Key: "url", Value: url}}
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

// GetURLFromHash returns URL given hash
func (m *MongoURLStore) GetURLFromHash(hash string) string {
	urlHashPair := URLHashPair{}
	filter := bson.D{{Key: "hash", Value: hash}}
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

	client, cancel, ctx := config.Connect()
	defer cancel()
	defer client.Disconnect(ctx)
	server := NewMongoServer(client)

	fmt.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port %s %v", port, err)
	}

}
