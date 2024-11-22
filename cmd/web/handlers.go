package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

// Home Handler
func (log *applog) home(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		// If there's an error, set a default value
		hostname = "unknown"
	}

	w.Header().Add("server", hostname)

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.serverError(w, r, err)
		return
	}

	err = t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.serverError(w, r, err)
		return
	}

}

// View Snnipet Handeler
func (log *applog) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with id %d", id)
}

// Snnipet Create
func (log *applog) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Snippet..."))
}

func (log *applog) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		// If there's an error, set a default value
		hostname = "unknown"
	}

	w.Header().Add("server", hostname)

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new Snippet..."))
}
