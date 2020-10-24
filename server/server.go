package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const HtmlContentType = "text/html; charset=UTF-8"
const ApplicationJSON = "application/json"
const ClientRoute = "../client/index.html"

type HashStore interface {
	GetHashFromURL(url string) string
	GetURLFromHash(hash string) string
	HashURL(url string) string
}

// HashingServer implements an instance of the server
type HashingServer struct {
	store HashStore
}

func (h *HashingServer) processHashing(w http.ResponseWriter, r *http.Request) {
	var requestBody struct{ URL string }
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	hash := h.store.GetHashFromURL(requestBody.URL)

	body := URLHashPair{
		URL:  requestBody.URL,
		Hash: hash,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("had trouble converting {URL: %s, hash: %s} to JSON, %v", requestBody.URL, hash, err)
	}

	w.Header().Set("content-type", ApplicationJSON)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonBody))
}

func (h *HashingServer) serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", HtmlContentType)
	http.ServeFile(w, r, ClientRoute)
}
func (h *HashingServer) checkHashAndRedirect(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Path[1:]
	URL := h.store.GetURLFromHash(hash)
	if URL == "" {
		fmt.Println("redirect to 404")
		http.Redirect(w, r, "/404", http.StatusNotFound)
	} else {
		fmt.Printf("redirect to %s", URL)
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

// RedirectServer implements the server that handles redirecting
func (h *HashingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := http.NewServeMux()
	router.Handle("/api/create-hash", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.processHashing(w, r)
	}))

	router.Handle("/404", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		h.serveHome(w, r)
	}))

	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/" {
			h.serveHome(w, r)
		} else {
			h.checkHashAndRedirect(w, r)
		}
	}))

	router.ServeHTTP(w, r)
}
