package main

import (
	"fmt"      // for printing or returning formated strings
	"log"      // for sending error logs
	"net/http" // for creating http servers and routes

	"github.com/SarakshiKaur/Go-Movie-Project/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize Gorilla mux
	r := mux.NewRouter()

	// Using mux to handle route
	// Here we can also tell the request method directly
	r.HandleFunc("/", handler.HandleRoot).Methods("GET")
	r.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", handler.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movie", handler.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", handler.UpdateMovie).Methods("PUT")

	fmt.Println("Started server on port 3000")

	// we are passing mux's souter instad of nil so that Gorilla muz handles the routing
	// for us instead of default http handler
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Server failed to run properly err: %v", err)
	}
}
