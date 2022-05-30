package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getMovies(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range movies {
		if item.ID == params["id"] {
			return json.NewEncoder(res).Encode(item)
		}
	}

	http.Error(res, "Movie not found", http.StatusNotFound)
	return nil
}

func createMovie(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(req.Body).Decode(&movie)
	if err != nil {
		http.Error(res, "Error decoding data", http.StatusInternalServerError)
		return err
	}

	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	return json.NewEncoder(res).Encode(movie)
}

func updateMovie(res http.ResponseWriter, req *http.Request) error {
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
		return nil
	}

	var movie Movie
	err := json.NewDecoder(req.Body).Decode(&movie)
	if err != nil {
		return err
	}

	movie.ID = params["id"]
	movies[movieIndex] = movie

	return json.NewEncoder(res).Encode(movie)
}

func deleteMovie(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	return json.NewEncoder(res).Encode(movies)
}
