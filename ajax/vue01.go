package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Pelicula struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Rating   int    `json:"rating"`
	Image    string `json:"url"`
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "vue01.html")
}

func peliculas(w http.ResponseWriter, r *http.Request) {
	peliculas := []Pelicula{
		{
			Name:     "The Avengers",
			Year:     2012,
			Director: "Joss Whedon",
			Rating:   5,
			Image:    "https://generacionxbox.com/wp-content/uploads/2019/05/avengers-xboxone-marvel-generacion-xbox-940x529.jpg",
		},
		{
			Name:     "Star Wars: Episode III",
			Year:     2005,
			Director: "George Lucas",
			Rating:   5,
			Image:    "https://static.wikia.nocookie.net/cine/images/2/23/Star-Wars-Episode-III-Revenge-of-the-Sith.jpg/revision/latest/top-crop/width/360/height/450?cb=20120930201011",
		},
		{
			Name:     "The Lord of the Rings: The Return of the King",
			Year:     2003,
			Director: "Peter Jackson",
			Rating:   5,
			Image:    "https://upload.wikimedia.org/wikipedia/en/b/be/The_Lord_of_the_Rings_-_The_Return_of_the_King_%282003%29.jpg",
		},
		{
			Name:     "The Fast and the Furious",
			Year:     2001,
			Director: "Rob Cohen",
			Rating:   5,
			Image:    "https://static.wikia.nocookie.net/doblaje/images/4/47/RFCartelOriginal.jpg/revision/latest?cb=20171125061739&path-prefix=es",
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
