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

	"github.com/facutk/golaburo/routes"
	"github.com/facutk/golaburo/utils"
)

var dbFile = "/data/foo.db"
var bucket = utils.GetEnv("LITESTREAM_BUCKET", "")

func main() {
	log.Println("main: init")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	lsdb, _ := utils.Litestream(ctx, dbFile, bucket)
	defer lsdb.SoftClose()

	utils.ConnectRedis()
	utils.InitTokenAuth()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("golang"))
	})

	r.Get("/uppercase/{str}", routes.GetUppercaseHandler)
	r.Get("/ping", routes.GetPingHandler)
	r.Post("/note", routes.NotePostHandler)
	r.Get("/note", routes.NoteGetHandler)
	r.Get("/sqlite", routes.GetSqliteSchemaVersion)
	r.Get("/v", routes.GetMigrationVersion)
	r.Get("/hits", routes.GetHits)
	r.Get("/redis", routes.GetRedisHandler)
	r.Get("/redis/u", routes.GetRedisUpdateHandler)
	r.Get("/login", routes.GetLoginHandler)
	r.Get("/env", routes.GetEnv)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/admin", routes.GetProtectedHandler)
	})

	http.ListenAndServe(":8080", r)

	<-ctx.Done()
	log.Print("gecko: received signal, shutting down")
}
