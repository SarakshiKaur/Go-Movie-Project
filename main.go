package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json" // for encoding and decoding things to json
	"fmt"           // for printing or returning formated strings
	"log"           // for sending error logs
	"math/rand"     // genrating random numbers
	"net/http"      // for creating http servers and routes
	"strconv"       // for converting between strings and numbers
	"strings"       // for string manipulation

	"github.com/gorilla/mux"
)

// this format of capital lettered names in struct
// is the standard for exported fields in Go to json
// small lettered names don't get exposed
// also after defining a propery and its type we specify that in json they will be reffered as
// the name we are giving now
// Title  string `json:"title"` means that Title property in go will be title when converted
// to json
type Movie struct {
	ID       string    `json:"id"`
	Imdb_id  string    `json:"imdb_id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type MovieRequest struct {
	Title    string    `json:"title"`
	Imdb_id  string    `json:"imdb_id"`
	Director *Director `json:"director"`
}

// slice or array of type Movie struct
var movies = []Movie{
	{
		ID:      "52420926a8d403",
		Imdb_id: "tt1375666",
		Title:   "Inception",
		Director: &Director{ // we are passing address of struct cause we want to modify actual struct
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:      "3fd26c41ae733f",
		Imdb_id: "tt0133093",
		Title:   "The Matrix",
		Director: &Director{
			Firstname: "Lana",
			Lastname:  "Wachowski",
		},
	},
	{
		ID:      "9074185900698c",
		Imdb_id: "tt0468569",
		Title:   "The Dark Knight",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:      "e2fccb318304cf",
		Imdb_id: "tt1285016",
		Title:   "The Social Network",
		Director: &Director{
			Firstname: "David",
			Lastname:  "Fincher",
		},
	},
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func checkIfIdExists(id string) (bool, int) {
	// we are checking if id exists in the movies array
	// _ is the index place we don't want that so we are ignoring it using _
	// movie variable will have actuall value of each movie from movies array
	for i, movie := range movies {
		if movie.ID == id {
			return true, i
		}
	}

	return false, 0
}

func exceptionHandler(newMovie MovieRequest) (bool, string) {
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

func checkIfMovieExists(newMovie MovieRequest) bool {
	// we are checking if movieName exists in the movies array
	movieName := newMovie.Title
	movieDirFirstName := newMovie.Director.Firstname
	movieDirLastname := newMovie.Director.Lastname

	for _, movie := range movies {
		if movie.Title == movieName &&
			movie.Director.Firstname == movieDirFirstName &&
			movie.Director.Lastname == movieDirLastname {
			return true
		}
	}
	return false
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	check, idx := checkIfIdExists(id)
	if !check {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	movie := movies[idx]
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	check, idx := checkIfIdExists(id)
	if !check {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	movieName := movies[idx].Title
	// In go if we want to slice a particular index we remove it like this
	// append means we are adding somehting to array
	// movies[:idx] will get all movies before that particular id (this will be source array in which we append elements)
	// movies[idx+1:] will get all movies after that particular id (this will give us array part we want to append on source one)
	// ... is a spread operator it opens up array and extract individual elements from it
	// eg => [4,5] will become 4,5 if we use ... operator
	// So whole array will be updated and only the element with that id will be removed
	movies = append(movies[:idx], movies[idx+1:]...)
	fmt.Fprintf(w, "Succesfully deleted movie: %v", movieName)
}

func generateID(movieName string) string {
	hash := sha256.Sum256([]byte(movieName))
	hashHex := hex.EncodeToString(hash[:]) // convert to hex string

	first6chars := hashHex[:6]
	last6chars := hashHex[len(hashHex)-6:]
	randomNum := strconv.Itoa(rand.Intn(100)) // generate's rand number between 0-99
	id := first6chars + randomNum + last6chars

	return id
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movieInfo MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieInfo)
	defer r.Body.Close() // close connection after reading body
	if err != nil {
		http.Error(w, "Incorrect data in request", http.StatusBadRequest)
		return
	}

	// handle eny empty values
	if check, str := exceptionHandler(movieInfo); check {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	if check := checkIfMovieExists(movieInfo); check {
		http.Error(w, "Movie with this Name already exists", http.StatusBadRequest)
		return
	}

	id := generateID(movieInfo.Title)

	newMovie := Movie{
		ID:      id,
		Imdb_id: movieInfo.Imdb_id,
		Title:   movieInfo.Title,
		Director: &Director{ // we are passing address of struct cause we want to modify actual struct
			Firstname: movieInfo.Director.Firstname,
			Lastname:  movieInfo.Director.Lastname,
		},
	}

	movies = append(movies, newMovie)
	fmt.Fprintf(w, "New Movie created with id: %v\n", id)
	json.NewEncoder(w).Encode(newMovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	check, idx := checkIfIdExists(id)
	if !check {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	// open up body json
	var movieInfo MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieInfo)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Incorrect data in request", http.StatusBadRequest)
		return
	}

	if check := checkIfMovieExists(movieInfo); check {
		http.Error(w, "Movie with this Name already exists, you can't update with same values", http.StatusBadRequest)
		return
	}

	// handle eny empty values
	if check, str := exceptionHandler(movieInfo); check {
		http.Error(w, str, http.StatusBadRequest)
		return
	}

	movies[idx] = Movie{
		ID:      id,
		Imdb_id: movieInfo.Imdb_id,
		Title:   movieInfo.Title,
		Director: &Director{
			Firstname: movieInfo.Director.Firstname,
			Lastname:  movieInfo.Director.Lastname,
		},
	}

	fmt.Fprintln(w, "Successfully updated the movie info")
	json.NewEncoder(w).Encode(movies[idx])
}

func main() {
	// Initialize Gorilla mux
	r := mux.NewRouter()

	// Using mux to handle route
	// Here we can also tell the request method directly
	r.HandleFunc("/", HandleRoot).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")

	fmt.Println("Started server on port 3000")

	// we are passing mux's souter instad of nil so that Gorilla muz handles the routing
	// for us instead of default http handler
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server failed to run properly err: %v", err)
	}
}
