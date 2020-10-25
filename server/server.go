package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-trimmer/utils"

	"github.com/gorilla/mux"
)

const HtmlContentType = "text/html; charset=utf-8"
const ApplicationJSON = "application/json"

const ClientRoute = "../client/build/index.html"

type HashStore interface {
	GetHashFromURL(url string) string
	GetURLFromHash(hash string) string
	HashURL(url string) string
}

// HashingServer implements an instance of the server
type HashingServer struct {
	store HashStore
}

func (h *HashingServer) createHashHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct{ URL string }
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Couldn't convert request body to JSON", http.StatusBadRequest)
		return
	}
	sanitizedURL, urlErr := utils.MakeURLAbsolute(requestBody.URL)
	if urlErr != nil {
		http.Error(w, "Bad URL", http.StatusBadRequest)
	}

	hash := h.store.GetHashFromURL(sanitizedURL)
	body := URLHashPair{
		URL:  sanitizedURL,
		Hash: hash,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "Had trouble hashing", http.StatusInternalServerError)
	}

	w.Header().Set("content-type", ApplicationJSON)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonBody))
}

func (h *HashingServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", HtmlContentType)
	http.ServeFile(w, r, ClientRoute)
}
func (h *HashingServer) redirectHandler(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	URL := h.store.GetURLFromHash(hash)
	if URL == "" {
		h.homeHandler(w, r)
	} else {
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

// RedirectServer implements the server that handles redirecting
func (h *HashingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	router.HandleFunc("/api/create-hash", h.createHashHandler)
	router.HandleFunc("/{hash}", h.redirectHandler)

	buildHandler := http.FileServer(http.Dir("../client/build"))
	router.PathPrefix("/").Handler(buildHandler)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("../client/build/static")))
	router.PathPrefix("/static/").Handler(staticHandler)

	router.ServeHTTP(w, r)
}
