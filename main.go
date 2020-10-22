package main

import (
	"log"
	"net/http"
)

const port = ":5000"

// InMemoryURLStore is temporary implementation
type InMemoryURLStore struct {
	hashes map[string]string
}

// HashURL hashes a given url
func (i *InMemoryURLStore) HashURL(url string) string {
	return "abc"
}

// GetHashFromURL is temporary implementation
func (i *InMemoryURLStore) GetHashFromURL(url string) string {
	if hash, ok := i.hashes[url]; ok {
		return hash
	}
	newHash := i.HashURL(url)
	i.hashes[url] = newHash
	return newHash
}

func (i *InMemoryURLStore) GetURLFromHash(url string) string {
	return ""
}

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://www.google.com", 301)
}

func main() {
	server := &HashingServer{&InMemoryURLStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
