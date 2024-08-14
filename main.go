package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

// get all
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// get by id
func getmovie_id(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

//create movie

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// delete Movie

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

//update movie

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)

		}
	}
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		Id:       "1",
		Title:    "Movie 1",
		Director: &Director{FirstName: "John", LastName: "Doe"},
	})
	movies = append(movies, Movie{
		Id:       "2",
		Title:    "Movie 2",
		Director: &Director{FirstName: "Jane", LastName: "Doe"},
	})

	//get all router
	r.HandleFunc("/movie", getMovies).Methods("GET")

	//get by id
	r.HandleFunc("/movie/{id}", getmovie_id).Methods("GET")

	//create
	r.HandleFunc("/create", createMovie).Methods("POST")

	//delete movie
	r.HandleFunc("/delete/{id}", deleteMovie).Methods("DELETE")

	//update movie
	r.HandleFunc("/update/{id}", updateMovie).Methods("PUT")

	fmt.Println("server started")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
