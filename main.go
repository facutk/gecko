package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"

	"github.com/facutk/golaburo/api"
	"github.com/facutk/golaburo/models"
	"github.com/facutk/golaburo/utils"
)

var dbFile = "/data/foo.db"
var bucket = "gecko-foo"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()
	log.Println("main: init")
	models.ConnectDB(dbFile)
	utils.Litestream(ctx, dbFile, bucket)
	utils.ConnectRedis()
	utils.InitTokenAuth()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("golang"))
	})

	r.Get("/uppercase/{str}", api.GetUppercaseHandler)
	r.Get("/ping", api.GetPingHandler)
	r.Post("/note", api.NotePostHandler)
	r.Get("/note", api.NoteGetHandler)
	r.Get("/sqlite", api.GetSqliteSchemaVersion)
	r.Get("/v", api.GetMigrationVersion)
	r.Get("/hits", api.GetHits)
	r.Get("/redis", api.GetRedisHandler)
	r.Get("/redis/u", api.GetRedisUpdateHandler)
	r.Get("/login", api.GetLoginHandler)
	r.Get("/env", api.GetEnv)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/admin", api.GetProtectedHandler)
	})

	http.ListenAndServe(":8080", r)

	<-ctx.Done()
	log.Print("gecko: received signal, shutting down")
}
