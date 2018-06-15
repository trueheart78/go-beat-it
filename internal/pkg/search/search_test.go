package search

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var game = "dark souls"

// TestSubmitValid tests the search.Submit() method
func TestSubmitValid(t *testing.T) {
	stub := stubValidSearch(game)
	html, err := Submit(game, stub.URL)

	assert.NotNil(t, html)
	assert.Nil(t, err)
}

// TestSubmitInvalid tests the search.Submit() method
func TestSubmitInvalid(t *testing.T) {
	stub := stubInvalidSearch()
	_, err := Submit(game, stub.URL)

	assert.NotNil(t, err)
}

func TestFormData(t *testing.T) {
	received := formData(game)
	expected := "detail=&length_max=&length_min=&length_type=main&plat=&queryString=dark+souls&sortd=Normal+Order&sorthead=popular&t=games"
	assert.Equal(t, expected, received)
}

func stubValidSearch(game string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		queryString := r.FormValue("queryString")
		t := r.FormValue("t")
		if game == queryString && t == "games" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		w.Write([]byte("sample response"))
	}))
}

func stubInvalidSearch() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
	}))
}
