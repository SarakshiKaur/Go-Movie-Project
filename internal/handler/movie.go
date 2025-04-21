package handler

import (
	"encoding/json" // for encoding and decoding things to json
	"fmt"           // for printing or returning formated strings
	"net/http"      // for creating http servers and routes

	"github.com/SarakshiKaur/Go-Movie-Project/internal/model"
	"github.com/SarakshiKaur/Go-Movie-Project/internal/service"
	"github.com/gorilla/mux" // advance package for better http routes creation
)

// /
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// /movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Movies)
}

// /movie/{id}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	check, idx := service.CheckIfIdExists(id)
	if !check {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	movie := model.Movies[idx]
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

// /movie/{id}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	check, idx := service.CheckIfIdExists(id)
	if !check {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	movieName := model.Movies[idx].Title
	// In go if we want to slice a particular index we remove it like this
	// append means we are adding somehting to array
	// movies[:idx] will get all movies before that particular id (this will be source array in which we append elements)
	// movies[idx+1:] will get all movies after that particular id (this will give us array part we want to append on source one)
	// ... is a spread operator it opens up array and extract individual elements from it
	// eg => [4,5] will become 4,5 if we use ... operator
	// So whole array will be updated and only the element with that id will be removed
	model.Movies = append(model.Movies[:idx], model.Movies[idx+1:]...)
	fmt.Fprintf(w, "Succesfully deleted movie: %v", movieName)
}

// /movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movieInfo model.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieInfo)
	defer r.Body.Close() // close connection after reading body
	if err != nil {
		http.Error(w, "Incorrect data in request", http.StatusBadRequest)
		return
	}

	// handle eny empty values
	if check, str := service.ExceptionHandler(movieInfo); check {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	if check := service.CheckIfMovieExists(movieInfo); check {
		http.Error(w, "Movie with this Name already exists", http.StatusBadRequest)
		return
	}

	id := service.GenerateID(movieInfo.Title)

	newMovie := model.Movie{
		ID:      id,
		Imdb_id: movieInfo.Imdb_id,
		Title:   movieInfo.Title,
		Director: &model.Director{ // we are passing address of struct cause we want to modify actual struct
			Firstname: movieInfo.Director.Firstname,
			Lastname:  movieInfo.Director.Lastname,
		},
	}

	model.Movies = append(model.Movies, newMovie)
	w.Header().Set("content-Type", "application/json")
	fmt.Fprintf(w, "New Movie created with id: %v\n", id)
	json.NewEncoder(w).Encode(newMovie)
}

// /movie/{id}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	check, idx := service.CheckIfIdExists(id)
	if !check {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	// open up body json
	var movieInfo model.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieInfo)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Incorrect data in request", http.StatusBadRequest)
		return
	}

	if check := service.CheckIfMovieExists(movieInfo); check {
		http.Error(w, "Movie with this Name already exists, you can't update with same values", http.StatusBadRequest)
		return
	}

	// handle eny empty values
	if check, str := service.ExceptionHandler(movieInfo); check {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	model.Movies[idx] = model.Movie{
		ID:      id,
		Imdb_id: movieInfo.Imdb_id,
		Title:   movieInfo.Title,
		Director: &model.Director{
			Firstname: movieInfo.Director.Firstname,
			Lastname:  movieInfo.Director.Lastname,
		},
	}

	w.Header().Set("content-Type", "application/json")
	fmt.Fprintln(w, "Successfully updated the movie info")
	json.NewEncoder(w).Encode(model.Movies[idx])
}
