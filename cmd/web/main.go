package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("Port is not set")
        PORT = "4000"
    }

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static",neuter(fileServer)))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id...}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Starting server on Port %s\n", PORT)

	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal("Error in Server Starting:", err)
}
