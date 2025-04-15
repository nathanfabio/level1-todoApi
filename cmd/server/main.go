package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func main() {
	r := chi.NewRouter()

	//Test
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})


	log.Println("Starting server on :7070")
	http.ListenAndServe(":7070", r)
}