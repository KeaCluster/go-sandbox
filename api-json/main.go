package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")

	http.ListenAndServe(":8000", r)
}
