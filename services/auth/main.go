package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("auth: init")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth"))
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login"))
	})

	http.ListenAndServe(":8080", r)

	log.Print("auth: received signal, shutting down")
}
