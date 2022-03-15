package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("golang"))
	})

	r.Get("/uppercase/{str}", getUppercaseHandler)
	r.Get("/hits", getHitsHandler)

	http.ListenAndServe(":8080", r)
}

func getUppercaseHandler(w http.ResponseWriter, r *http.Request) {
	str := chi.URLParam(r, "str")
	upperStr := strings.ToUpper(str)
	w.Write([]byte(upperStr))
}

func getHitsHandler(w http.ResponseWriter, r *http.Request) {
	hits := fmt.Sprintf("%d", 0)
	w.Write([]byte(hits))
}
