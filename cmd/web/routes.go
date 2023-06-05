package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	// Routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// Static route
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return secureHeaders(mux.ServeHTTP)

}
