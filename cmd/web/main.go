package main

import (
    "log"
    "net/http"
)

func main() {
	const PORT = "4000"

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id...}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Starting server on Port %s\n", PORT)

	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal("Error in Server Starting:", err)
}
