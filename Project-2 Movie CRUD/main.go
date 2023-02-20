package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movies struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movies

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", allMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", movie).Methods("GET")
	r.HandleFunc("/movie", createMovies).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovies).Methods("DELETE")
	fmt.Printf("Listeing from the port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Handler
func allMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	for _, item := range movies {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	rand.Seed(time.Now().UnixNano())
	movie.ID = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movies
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = param["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}

}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.ID == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}
