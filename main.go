package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stinkyfingers/catchphraseAPI/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s, err := handlers.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HandleStatus)
	mux.HandleFunc("/all", s.Cors(s.HandleAll))
	mux.HandleFunc("/upload", s.Cors(s.HandleUpload))

	log.Print("Running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
