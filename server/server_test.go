package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-trimmer/server/utils"
)

type StubHashStore struct {
	hashes []URLHashPair
}

func (s *StubHashStore) HashURL(url string) string {
	return "abc"
}

func (s *StubHashStore) GetHashFromURL(url string) string {
	for _, urlHashPair := range s.hashes {
		if urlHashPair.URL == url {
			return urlHashPair.Hash
		}
	}
	newHash := s.HashURL(url)
	s.hashes = append(s.hashes, URLHashPair{
		URL:  url,
		Hash: newHash,
	})
	return newHash
}

func (s *StubHashStore) GetURLFromHash(hash string) string {
	for _, urlHashPair := range s.hashes {
		if urlHashPair.Hash == hash {
			return urlHashPair.URL
		}
	}
	return ""
}

func TestHashingServer(t *testing.T) {
	store := StubHashStore{
		hashes: []URLHashPair{
			URLHashPair{
				Hash: "xyz",
				URL:  "https://google.com",
			},
		},
	}
	server := &HashingServer{&store}

	t.Run("returns the hash for a given url if it exist in hash map", func(t *testing.T) {
		url := "https://google.com"

		request := newHashingRequest(url)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		body := readHashingResponseBody(response)

		utils.AssertString(t, body.Hash, "xyz")
		assertStatus(t, response.Code, http.StatusCreated)
	})
	t.Run("returns a new hash for a given url if it doesn't exist in the hash map", func(t *testing.T) {
		url := "xyz.com"

		request := newHashingRequest(url)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		body := readHashingResponseBody(response)

		utils.AssertString(t, body.Hash, "abc")
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
	t.Run("serves react app, if hash doesn't exist on hash table ", func(t *testing.T) {
		hash := "nonsense-hash"
		request := newRedirectRequest(hash)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, HtmlContentType)
	})
	t.Run("serves react app on route '/' ", func(t *testing.T) {
		request := newHomePageRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, HtmlContentType)
	})
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
		t.Errorf("response did not have corrent content-type, got: %s, want: %s", want, response.Result().Header.Get("content-type"))
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
	req, _ := http.NewRequest(http.MethodPost, "/api/create-hash", bytes.NewBuffer(body))
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

func readHashingResponseBody(r *httptest.ResponseRecorder) URLHashPair {
	body := URLHashPair{}
	json.NewDecoder(r.Body).Decode(&body)
	return body
}
