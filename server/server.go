package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-trimmer/utils"
)

const HtmlContentType = "text/html; charset=UTF-8"
const ApplicationJSON = "application/json"
const TextCSS = "text/css"
const TextJavascript = "text/javascript; charset=UTF-8"

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

func (h *HashingServer) processHashing(w http.ResponseWriter, r *http.Request) {
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

func (h *HashingServer) serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", HtmlContentType)
	http.ServeFile(w, r, ClientRoute)
}
func (h *HashingServer) checkHashAndRedirect(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Path[1:]
	URL := h.store.GetURLFromHash(hash)
	if URL == "" {
		w.WriteHeader(http.StatusNotFound)
		h.serveHome(w, r)
	} else {
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

	router.Handle("/static/css/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../client/build" + r.URL.String()
		w.Header().Set("content-type", TextCSS)
		http.ServeFile(w, r, path)
	}))
	router.Handle("/static/js/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := "../client/build" + r.URL.String()
		w.Header().Set("content-type", TextJavascript)
		http.ServeFile(w, r, path)
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
