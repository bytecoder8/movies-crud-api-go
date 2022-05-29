package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
