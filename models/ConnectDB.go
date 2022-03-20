package models

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func ConnectDB() {
	log.Println("db: connecting")
	_db, err := sql.Open("sqlite", "/data/foo.db")
	if err != nil {
		log.Println("db: error", err)
		panic(err)
	}

	db = _db
	log.Println("db: connected")

	EnsureSchema(db)
}
