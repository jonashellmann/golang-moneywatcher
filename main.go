package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./a/")
	staticFileHandler := http.StripPrefix("/a/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/a/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":5656", r)
}
