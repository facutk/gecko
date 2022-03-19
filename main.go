package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/facutk/golaburo/api"
	"github.com/facutk/golaburo/models"
)

func main() {
	log.Println("main: init")
	models.ConnectDB()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("golang"))
	})

	r.Get("/uppercase/{str}", getUppercaseHandler)
	r.Get("/ping", api.GetPingHandler)
	r.Post("/note", api.NotePostHandler)
	r.Get("/note", api.NoteGetHandler)
	r.Get("/sqlite", api.GetSqliteSchemaVersion)
	r.Get("/v", api.GetMigrationVersion)
	r.Get("/hits", api.GetHits)

	http.ListenAndServe(":8080", r)
}

func getUppercaseHandler(w http.ResponseWriter, r *http.Request) {
	str := chi.URLParam(r, "str")
	upperStr := strings.ToUpper(str)
	w.Write([]byte(upperStr))
}
