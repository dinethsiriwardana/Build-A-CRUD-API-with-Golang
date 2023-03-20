package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Append All movies expt selected one
			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}

	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

// because its a json we cant just update it so we need to delete it and reinsert
func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Append All movies expt selected one
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "738227", Title: "Movie One", Director: &Director{Firstname: "Jhon", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "895727", Title: "Movie Two", Director: &Director{Firstname: "Mike", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "895727", Title: "Movie Three", Director: &Director{Firstname: "Ronald", Lastname: "Potter"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET") // getMovies -> bulk / getMovie -> Single
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Stating server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
