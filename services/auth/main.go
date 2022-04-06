package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("auth: hello golang x")

	r := chi.NewRouter()

	r.Get("/hits", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
		if err != nil {
			log.Println("db: error", err)
			panic(err)
		}

		row := db.QueryRow("select hits from visits where id=1;")

		var hits int
		row.Scan(&hits)
		log.Println("auth: hits: ", hits)

		_, err2 := db.Exec("update visits set hits=$1 where id=1", hits+1)
		if err2 != nil {
			log.Println("auth: hits update failed 1", err)
		}

		fmt.Fprint(w, hits)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "authx ss")
	})

	http.ListenAndServe(":8080", r)
}
