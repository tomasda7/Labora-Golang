package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tomasda7/movies_api/models"
	"github.com/tomasda7/movies_api/services"
)

// Create
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie models.Movie

	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMovieName, err := services.CreateMovie(newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("An error occurred while trying to create a new movie."))
		return
	}

	res := fmt.Sprintf("The movie '%s' was created successfully!", newMovieName)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// Read
func GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := services.GetAllMovies()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("An error occurred while trying to retrieve the movies."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

// Update
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// id by params
	paramId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(paramId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		msg := fmt.Sprintf("Cannot parse id from '%s.", paramId)
		w.Write([]byte(msg))
		return
	}

	// validate that the movie exists
	err = MovieExists(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update movie data
	var updMovie models.Movie

	err = json.NewDecoder(r.Body).Decode(&updMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updMovieID, err := services.UpdateMovie(id, updMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("An error occurred while trying to update the movie."))
		return
	}

	// response
	res := fmt.Sprintf("The movie with the id '%d' was updated successfully!", updMovieID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// Delete
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	// id by params
	paramId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(paramId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		msg := fmt.Sprintf("Cannot parse id from '%s.", paramId)
		w.Write([]byte(msg))
		return
	}

	// validate that the movie exists
	err = MovieExists(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// delete movie
	movieName, err := services.DeleteMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("An error occurred while trying to delete a movie."))
		return
	}
	// response
	res := fmt.Sprintf("The movie '%s' was deleted successfully.", movieName)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
