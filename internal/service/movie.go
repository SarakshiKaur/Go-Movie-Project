package service

import (
	"strings" // for string manipulation

	"github.com/SarakshiKaur/Go-Movie-Project/internal/model"
)

func CheckIfMovieExists(newMovie model.MovieRequest) bool {
	// we are checking if movieName exists in the movies array
	movieName := newMovie.Title
	movieDirFirstName := newMovie.Director.Firstname
	movieDirLastname := newMovie.Director.Lastname

	for _, movie := range model.Movies {
		if movie.Title == movieName &&
			movie.Director.Firstname == movieDirFirstName &&
			movie.Director.Lastname == movieDirLastname {
			return true
		}
	}
	return false
}

func ExceptionHandler(newMovie model.MovieRequest) (bool, string) {
	// trimspace removes extra spaces
	if strings.TrimSpace(newMovie.Title) == "" {
		return true, "Title cannot be empty"
	}
	if strings.TrimSpace(newMovie.Imdb_id) == "" {
		return true, "Imdb_id cannot be empty"
	}
	if strings.TrimSpace(newMovie.Director.Firstname) == "" {
		return true, "Firstname cannot be empty"
	}
	if strings.TrimSpace(newMovie.Director.Lastname) == "" {
		return true, "Lastname cannot be empty"
	}

	return false, ""
}

func CheckIfIdExists(id string) (bool, int) {
	// we are checking if id exists in the movies array
	// _ is the index place we don't want that so we are ignoring it using _
	// movie variable will have actuall value of each movie from movies array
	for i, movie := range model.Movies {
		if movie.ID == id {
			return true, i
		}
	}

	return false, 0
}
