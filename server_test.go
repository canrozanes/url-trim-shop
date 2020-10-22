package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubHashStore struct {
	hashes map[string]string
}

func (s *StubHashStore) HashURL(url string) string {
	return "abc"
}

func (s *StubHashStore) GetHashFromURL(url string) string {
	for key, val := range s.hashes {
		if val == url {
			return key
		}
	}
	newHash := s.HashURL(url)
	s.hashes[newHash] = url
	return newHash
}

func (s *StubHashStore) GetURLFromHash(hash string) string {
	if url, ok := s.hashes[hash]; ok {
		return url
	}
	return ""
}

func TestHashingServer(t *testing.T) {
	store := StubHashStore{
		hashes: map[string]string{
			"xyz": "https://google.com",
		},
	}
	server := &HashingServer{&store}

	t.Run("returns the hash for a given url if it exist in hash map", func(t *testing.T) {
		url := "https://google.com"

		request := newHashingRequest(url)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertString(t, response.Body.String(), "xyz")
		assertStatus(t, response.Code, http.StatusCreated)
	})
	t.Run("returns a new hash for a given url if it doesn't exist in the hash map", func(t *testing.T) {
		url := "xyz.com"

		request := newHashingRequest(url)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertString(t, response.Body.String(), "abc")
		assertStatus(t, response.Code, http.StatusCreated)
	})
	t.Run("redirects to route '/[hash]', if hash exists on hash table", func(t *testing.T) {
		hash := "xyz"
		request := newRedirectRequest(hash)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusFound)
		assertRedirectURL(t, response, "https://google.com")
	})
	t.Run("redirects to route '/404', if hash doesn't exist on hash table ", func(t *testing.T) {
		hash := "nonsense-hash"
		request := newRedirectRequest(hash)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusFound)
		assertRedirectURL(t, response, "/404")
	})
	t.Run("returns html on route '/' ", func(t *testing.T) {
		request := newHomePageRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, HtmlContentType)
	})
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func assertRedirectURL(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := response.HeaderMap.Get("Location")
	if got != want {
		t.Errorf("redirect url got %s, want %v", got, want)
	}
}

func newHashingRequest(url string) *http.Request {
	body := createRequestBody(url)
	req, _ := http.NewRequest(http.MethodPost, "/create-hash", bytes.NewBuffer(body))
	return req
}

func newHomePageRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	return req
}

func newRedirectRequest(hash string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/"+hash, nil)
	return req
}

func createRequestBody(url string) []byte {
	body := struct {
		URL string
	}{
		URL: url,
	}
	jsonBody, _ := json.Marshal(body)

	return jsonBody
}