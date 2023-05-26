package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Command line flags
	addr := flag.String("addr", ":5000", "HTTP network address")
	staticDir := flag.String("staticDir", "./ui/static", "Path to folder with static files")
	flag.Parse()

	// Static route
	fileServer := http.FileServer(http.Dir(*staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Listen
	log.Println("Starting server on", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
