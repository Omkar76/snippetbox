package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Hello from Snippetbox"))
	files := []string{
		"./ui/html/pages/base.go.tpl",
		"./ui/html/pages/home.go.tpl",
		"./ui/html/partials/nav.go.tpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if ts.ExecuteTemplate(w, "base", nil) != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", http.MethodPost)

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
