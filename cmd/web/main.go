package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)
	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
