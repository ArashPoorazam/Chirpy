package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	fileServerHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))
	cfg := &apiConfig{}

	// handlers
	mux.Handle("/app/", cfg.middlewareCount(fileServerHandler))
	mux.HandleFunc("GET /admin/metrics", cfg.metricHits)
	mux.HandleFunc("POST /api/reset", cfg.resetHits)
	mux.HandleFunc("GET /api/healthz", handleHealthz)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// ListenAndServe
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
