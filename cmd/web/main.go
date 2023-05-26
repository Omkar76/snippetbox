package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Command line flags
	addr := flag.String("addr", ":5000", "HTTP network address")
	staticDir := flag.String("staticDir", "./ui/static", "Path to folder with static files")
	flag.Parse()

	// Loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Static route
	fileServer := http.FileServer(http.Dir(*staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Listen
	infoLog.Println("Starting server on", *addr)
	// err := http.ListenAndServe(*addr, mux)

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
