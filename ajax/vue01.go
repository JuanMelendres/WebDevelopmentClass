package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Pelicula struct {
	Name string `json:"name"`
	Year int `json:"year"`
	Director string `json:"director"`
	Rating int `json:"rating"`
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "vue01.html")
}

func peliculas(w http.ResponseWriter, r *http.Request) {
	peliculas := []Pelicula{
		{
			Name: "The Avengers",
			Year: 2012,
			Director: "Joss Whedon",
			Rating: 5,
		},
		{
			Name: "Star Wars: Episode III",
			Year: 2005,
			Director: "George Lucas",
			Rating: 5,
		},
		{
			Name: "The Lord of the Rings: The Return of the King",
			Year: 2003,
			Director: "Peter Jackson",
			Rating: 5,
		},
		{
			Name: "The Fast and the Furious",
			Year: 2001,
			Director: "Rob Cohen",
			Rating: 5,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peliculas)
}

func main() {
	http.HandleFunc("/movie", peliculas)
	http.HandleFunc("/", mostrarHTML)

	log.Printf("Starting server on localhost:8080...\n")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}