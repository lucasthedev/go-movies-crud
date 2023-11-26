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
	Id       string    `json:"id"`
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
	fmt.Println(movies)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	routes := mux.NewRouter()

	movies = append(movies, Movie{
		Id:    "1",
		Isbn:  "876432",
		Title: "Harry Potter",
		Director: &Director{
			Firstname: "Lucas",
			Lastname:  "Pereira",
		},
	})

	movies = append(movies, Movie{
		Id:    "2",
		Isbn:  "597142",
		Title: "Veloz e Furiozos",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	routes.HandleFunc("/movies", getMovies).Methods("GET")
	routes.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	routes.HandleFunc("/movies", createMovie).Methods("POST")
	routes.HandleFunc("/movies/{id}", updateMovie).Methods("UPDATE")
	routes.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", routes))
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-type", "application/json")
	var movie Movie
	params := mux.Vars(r)
	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = strconv.Itoa(rand.Intn(1000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			break
		}
	}

}
