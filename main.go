package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()

	// handlers
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// ListenAndServe
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
