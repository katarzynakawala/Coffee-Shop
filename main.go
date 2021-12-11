package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	w.Write([]byte("Hello from Coffee Shop"))
}

func displayCoffee(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a coffee"))
}

func createCoffee(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new coffee"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/coffee", displayCoffee)
	mux.HandleFunc("/coffee/create", createCoffee)

	log.Println("Starting server in: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}