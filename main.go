package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

const PORT = 8080

func createRouter() *mux.Router {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PATCH")

	return r
}

func init() {
	movies = append(movies, Movie{
		ID: "1", Isbn: "23456", Title: "First Movie",
		Director: &Director{
			FirstName: "Steve",
			LastName:  "Jobs",
		},
	})

	movies = append(movies, Movie{
		ID: "2", Isbn: "564212", Title: "Second Movie",
		Director: &Director{
			FirstName: "Garry",
			LastName:  "Smith",
		},
	})
}

func main() {
	router := createRouter()

	fmt.Printf("Starting server at PORT:%d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), router); err != nil {
		log.Fatal(err)
	}
}
