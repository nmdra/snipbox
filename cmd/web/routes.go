package main

import (
	"net/http"
)

func (log *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", neuter(fileServer)))

	mux.HandleFunc("GET /{$}", log.home)
	mux.HandleFunc("GET /snippet/view/{id...}", log.snippetView)
	mux.HandleFunc("GET /snippet/create", log.snippetCreate)

	mux.HandleFunc("POST /snippet/create", log.snippetCreatePost)

	return mux
}
