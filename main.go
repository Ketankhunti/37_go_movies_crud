package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
);

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director Director `json:"director"`
}

type Director struct{
   Firstname string `json:"firstname"`
   Lastname string `json:"lastname"`
   
}

var movies []Movie


func getMovies(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json");
	json.NewEncoder(w).Encode(movies);
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json");
	 vars := mux.Vars(r);
     var found bool = false;
	 movie := Movie{};
	 for index,v := range movies {
		if v.ID == vars["id"] {
			movie = v;
			movies = append(movies[:index],movies[index+1:]...);
			found = true;
			break;
		}
	 }
	  json.NewEncoder(w).Encode(movie);
	 if found {
		w.Write([]byte("movie deleted successfully"))
	 }else{
		w.Write([]byte("could not find movie"));
	 }
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json");

	param := mux.Vars(r);

	for _,v := range movies {
		if v.ID == param["id"] {
			json.NewEncoder(w).Encode(v);
			return;
		}
	}
	json.NewEncoder(w).Encode("movie did not found");
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json");

	movie := Movie{};

	json.NewDecoder(r.Body).Decode(&movie);

	movie.ID = strconv.Itoa(rand.Intn(1000000)) ;

	movies = append(movies,movie);

	json.NewEncoder(w).Encode(movie);
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type","application/json");
	 movie := Movie{};
	 json.NewDecoder(r.Body).Decode(&movie);
	 params := mux.Vars(r);
	 for index,v := range movies {
		if v.ID == params["id"] {
           movies = append(movies[:index],movies[index+1:]...);
		   movie.ID = params["id"];
		   movies = append(movies,movie);
		   json.NewEncoder(w).Encode(movie);
		}
	 }
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to ketan's API</h1>"));
}

func main() {
	r := mux.NewRouter();

	movies = append(movies, Movie{ID:"1",Isbn: "123456",Title:"Movie One",Director: Director{Firstname:"ketan",Lastname:"khunti"}});
	movies = append(movies, Movie{ID:"2",Isbn: "345626",Title:"Movie TWo",Director:Director{Firstname:"karan",Lastname:"khunti"}});

	r.HandleFunc("/",hello).Methods("GET");
	r.HandleFunc("/movies",getMovies).Methods("GET");
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET");
	r.HandleFunc("/movies",createMovie).Methods("POST");
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT");
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE");

	fmt.Println("starting server at port : 4000");

	log.Fatal(http.ListenAndServe(":4000",r));


}
