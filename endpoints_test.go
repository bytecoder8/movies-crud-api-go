package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMovies(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := createRouter()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: %v want %v", status, http.StatusOK)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: %v want %v", ctype, "application/json")
	}
}

func TestGetMovie(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := createRouter()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json")
	}

	expected := `{"id":"1","isbn":"23456","title":"First Movie","director":{"firstname":"Steve","lastname":"Jobs"}}` + "\n"
	if body := rr.Body.String(); body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expected)
	}
}
