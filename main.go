package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func getIndex(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello! 🐶"))
}

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}

	http.Error(res, "Movie not found", http.StatusNotFound)
}

func createMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(req.Body).Decode(&movie)
	if err != nil {
		http.Error(res, "Error decoding data", http.StatusInternalServerError)
		return
	}

	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movie)
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	params := mux.Vars(req)

	movieIndex := -1
	for index, item := range movies {
		if item.ID == params["id"] {
			movieIndex = index
			break
		}
	}

	if movieIndex == -1 {
		http.Error(res, "Movie not found", http.StatusNotFound)
	}

	var movie Movie
	err := json.NewDecoder(req.Body).Decode(&movie)
	if err != nil {
		http.Error(res, "Error decoding data", http.StatusInternalServerError)
		return
	}

	movie.ID = params["id"]
	movies[movieIndex] = movie

	json.NewEncoder(res).Encode(movie)
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(res).Encode(movies)
}

const PORT = 8080

func main() {
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

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PATCH")

	fmt.Printf("Starting server at PORT:%d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), r); err != nil {
		log.Fatal(err)
	}
}
