package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stinkyfingers/catchphraseAPI/handlers"
)

func main() {
	port := "8080"
	if p := os.Getenv("port"); p != "" {
		port = p
	}

	s, err := handlers.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.Cors(s.HandleStatus))
	mux.HandleFunc("/all", s.Cors(s.HandleAll))
	mux.HandleFunc("/upload", s.Cors(s.HandleUpload))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
