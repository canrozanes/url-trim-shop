package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const HtmlContentType = "text/html; charset=UTF-8"
const ClientRoute = "./client/index.html"

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
	var response struct{ URL string }
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	hash := h.store.GetHashFromURL(response.URL)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, hash)
}

func (h *HashingServer) serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", HtmlContentType)
	http.ServeFile(w, r, ClientRoute)
}
func (h *HashingServer) checkHashAndRedirect(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Path[1:]
	URL := h.store.GetURLFromHash(hash)
	if URL == "" {
		http.Redirect(w, r, "/404", http.StatusFound)
	} else {
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

// RedirectServer implements the server that handles redirecting
func (h *HashingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := http.NewServeMux()
	router.Handle("/create-hash", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		h.processHashing(w, r)
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
