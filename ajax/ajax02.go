package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var a = [5]int{2, 4, 6, 8, 10}

type book struct {
	Titulo string `json: titulo`
	Autor  string `json: autor`
}

type user struct {
	UserId string `json: userId`
	ID  string `json: id`
  Title  string `json: title`
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ajax02.html")
}

func darMensaje(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(a)
	}
	if r.Method == "POST" {
		r.ParseForm()
		x := r.Form.Get("y")
		i1, err := strconv.ParseInt(x, 10, 64)
		if err == nil {
			json.NewEncoder(w).Encode(a[i1])
		}
	}
}

func darMensaje2(w http.ResponseWriter, r *http.Request) {
	// libro := book{
	// 	Titulo: "La Casa",
	// 	Autor:  "Paco Roca",
	// }
  data := user {
    UserId: "1",
    ID: "1",
    Title: "delectus aut autem",
  }
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {

	http.HandleFunc("/dato", darMensaje)
  http.HandleFunc("/json", darMensaje2)
	http.HandleFunc("/", mostrarHTML)

	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
