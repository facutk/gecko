package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/facutk/golaburo/api"
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
	r.Get("/ping", api.GetPingHandler)
	r.Get("/write/{str}", api.GetWriteHandler)
	r.Get("/read", api.GetReadHandler)

	http.ListenAndServe(":8080", r)
}

func getUppercaseHandler(w http.ResponseWriter, r *http.Request) {
	str := chi.URLParam(r, "str")
	upperStr := strings.ToUpper(str)
	w.Write([]byte(upperStr))
}

func getHitsHandler(w http.ResponseWriter, r *http.Request) {
	hits := fmt.Sprintf("%d", 1)
	w.Write([]byte(hits))
}
