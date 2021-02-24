package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/nojson", ServeNoJson)
	http.HandleFunc("/actividad", ServeSecondFile)
	http.HandleFunc("/data", ServeTestJson)
	http.HandleFunc("/", MainHtml)
	log.Printf("Starting server on localhost:8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MainHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func ServeNoJson(w http.ResponseWriter, r *http.Request) {
	log.Printf("request nojson file: %s", r.Host)
	http.ServeFile(w, r, "nojson.txt")
}

func ServeSecondFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("request secondfile: %s", r.Host)
	http.ServeFile(w, r, "actividad.txt")
}

func ServeTestJson(w http.ResponseWriter, r *http.Request) {
	log.Printf("request testjson: %s", r.Host)
	http.ServeFile(w, r, "data.json")
}

// JQuery
// es una biblioteca multiplataforma de JavaScript, creada inicialmente por John Resig,
// que permite simplificar la manera de interactuar con los documentos HTML, manipular el árbol DOM, manejar eventos,
// desarrollar animaciones y agregar interacción con la técnica AJAX a páginas web.

// Ventajas
// Nos permite realizar consultas sobre la estructura del DOM de la web y realizar personalizaciones a medida como
// aplicar diferentes estilos y efectos en función de los eventos que definamos.

// Desventajas
// La gran cantidad de versiones publicadas
// No importa si usted está corriendo la última versión de jQuery, usted tendrá que hostear la librería usted mismo,
//  o descargar la librería desde Google
