package main

import (
	"fmt"

	"net/http"
)

func darMensaje(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World, %s! \n", r)
	// fmt.Fprintf(w, "User info, %s! \n", r.UserAgent())
	// fmt.Fprintf(w, "Method, %s! \n", r.Method)
	// fmt.Fprintf(w, "url, %s! \n", r.URL)
  // fmt.Fprintf(w, "Host, %s! \n", r.Host) 
  // fmt.Fprintf(w, "Cookies, %s! \n", r.Cookies()) Location
  fmt.Fprintf(w, "RemoteAddr, %s! \n", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", darMensaje)
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
