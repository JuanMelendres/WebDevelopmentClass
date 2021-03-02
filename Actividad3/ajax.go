package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Data struct {
	mu   sync.Mutex
	data []int
}

type Value struct {
	Arr []int `json:"arr"`
}

var sd = Data{
	data: make([]int, 0),
}

func (d *Data) GetNewValue() []int {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	return sd.data
}

func (d *Data) AddNewValue(num int) []int {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	sd.data = append(sd.data, num)
	return sd.data
}

func actividad(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	count := r.Form.Get("counter")
	value, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println("Error")
	}
	new := Value{
		Arr: sd.GetNewValue()[value:],
	}
	log.Printf("New data sent: %v", new.Arr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(new)
}

func newData() {
	for {
		time.Sleep(4 * time.Second)
		new := sd.AddNewValue(rand.Intn(300))
		log.Printf("Array: %v", new)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./"))

	http.HandleFunc("/Actividad3", actividad)
	http.Handle("/", fs)

	go newData()

	log.Println("Starting server on localhost:8080...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}