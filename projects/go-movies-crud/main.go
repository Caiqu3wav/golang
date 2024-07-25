package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
		ID string `json:"id"`
		Isbn string `json:"isbn"`
		Title string `json:"title"`
		Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().set("content-type", "application/json")
	json.encoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "LyfeStyle", Director : &Director(
		FirstName: "Caique", 
	LastName: "Gonçalves")})

	movies = append(movies, Movie{ID: "2", Isbn: "468227", Title: "Pissed", Director : &Director(
		FirstName: "Caique", 
	LastName: "Gonçalves")})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.fatal(http.ListenAndServe(":8080", r))
}