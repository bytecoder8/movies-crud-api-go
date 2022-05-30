package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetMovies(t *testing.T) {
	seedData()
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
	seedData()
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

func TestDeleteMovie(t *testing.T) {
	seedData()
	req, err := http.NewRequest("DELETE", "/movies/1", nil)
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

	body := rr.Body.String()
	if strings.Contains(body, `"id":"1"`) {
		t.Error("Movie is still in list!")
	}
}

func TestCreateMovie(t *testing.T) {
	seedData()
	postData := `{
		"isbn" : "95473897",
		"title": "Third Movie",
		"director": {
			"firstname" : "New",
			"lastname" : "Director"
		}
	}`

	req, err := http.NewRequest("POST", "/movies", strings.NewReader(postData))
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

	expected := `{"id":"3","isbn":"95473897","title":"Third Movie","director":{"firstname":"New","lastname":"Director"}}` + "\n"
	if body := rr.Body.String(); body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestUpdateMovie(t *testing.T) {
	seedData()
	postData := `{
		"isbn" : "1111",
		"title": "Updated First Movie",
		"director": {
			"firstname" : "Steve",
			"lastname" : "Jobs"
		}
	}`

	req, err := http.NewRequest("PATCH", "/movies/1", strings.NewReader(postData))
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

	expected := `{"id":"1","isbn":"1111","title":"Updated First Movie","director":{"firstname":"Steve","lastname":"Jobs"}}` + "\n"
	if body := rr.Body.String(); body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expected)
	}
}
