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

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var movies []Movie

const PORT = 8080

func createRouter() *mux.Router {
	r := mux.NewRouter()

	// Routes
	r.Handle("/movies", appHandler(getMovies)).Methods("GET")
	r.Handle("/movies", appHandler(createMovie)).Methods("POST")
	r.Handle("/movies/{id}", appHandler(getMovie)).Methods("GET")
	r.Handle("/movies/{id}", appHandler(deleteMovie)).Methods("DELETE")
	r.Handle("/movies/{id}", appHandler(updateMovie)).Methods("PATCH")

	return r
}

func seedData() {
	movies = make([]Movie, 0)
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

func init() {
	seedData()
}

func main() {
	router := createRouter()

	fmt.Printf("Starting server at PORT:%d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), router); err != nil {
		log.Fatal(err)
	}
}
