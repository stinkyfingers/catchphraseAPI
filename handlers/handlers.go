package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stinkyfingers/catchphraseAPI/db"
	"github.com/stinkyfingers/catchphraseAPI/phrase"
	"gopkg.in/mgo.v2"
)

type Server struct {
	DB *mgo.Session
}

func NewServer() (*Server, error) {
	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}

	return &Server{
		DB: session,
	}, nil
}

func (s *Server) Cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "http://ocalhost:8081")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		h.ServeHTTP(w, r)
	}
}

func (s *Server) HandleStatus(w http.ResponseWriter, r *http.Request) {
	msg := "DB up"

	err := s.DB.Ping()
	if err != nil {
		msg = err.Error()
	}

	w.Write([]byte("API running \nDB Ping: " + msg))
}

func (s *Server) HandleAll(w http.ResponseWriter, r *http.Request) {
	categories, err := phrase.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (s *Server) HandleUpload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var cats []phrase.Category
	err = json.Unmarshal(b, &cats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, cat := range cats {
		temp := phrase.Category{Name: cat.Name}
		err = temp.Find()
		if err != nil && err != mgo.ErrNotFound {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if temp.ID.Valid() {
			err = temp.Remove()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		err = cat.Insert()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
