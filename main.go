package main

import (
	"encoding/json" // for encoding and decoding things to json
	"fmt"           // for printing or returning formated strings
	"log"           // for sending error logs
	"net/http"      // for creating http servers and routes
	"strconv"       // for converting strings

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
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// slice or array of type Movie struct
var movies = []Movie{
	{
		ID:    "0",
		Isbn:  "1234567890",
		Title: "Inception",
		Director: &Director{ // we are passing address of struct cause we want to modify actual struct
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:    "1",
		Isbn:  "9876543210",
		Title: "The Matrix",
		Director: &Director{
			Firstname: "Lana",
			Lastname:  "Wachowski",
		},
	},
	{
		ID:    "2",
		Isbn:  "1122334455",
		Title: "The Dark Knight",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:    "3",
		Isbn:  "2233445566",
		Title: "The Social Network",
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

func checkIfIdExists(id string) bool {
	// we are checking if id exists in the movies array
	// _ is the index place we don't want that so we are ignoring it using _
	// movie variable will have actuall value of each movie from movies array
	for _, movie := range movies {
		if movie.ID == id {
			return true
		}
	}

	return false
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !checkIfIdExists(vars["id"]) {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	movie := movies[id]
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !checkIfIdExists(vars["id"]) {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	movieName := movies[id].Title
	// In go if we want to slice a particular index we remove it like this
	// append means we are adding somehting to array
	// movies[:id] will get all movies before that particular id (this will be source array in which we append elements)
	// movies[id+1:] will get all movies after that particular id (this will give us array part we want to append on source one)
	// ... is a spread operator it opens up array and extract individual elements from it
	// eg => [4,5] will become 4,5 if we use ... operator
	// So whole array will be updated and only the element with that id will be removed
	movies = append(movies[:id], movies[id+1:]...)
	fmt.Fprintf(w, "Succesfully deleted movie: %v", movieName)
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

	fmt.Println("Started server on port 3000")

	// we are passing mux's souter instad of nil so that Gorilla muz handles the routing
	// for us instead of default http handler
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server failed to run properly err: %v", err)
	}
}
