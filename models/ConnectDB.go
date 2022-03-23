package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func ConnectDB(dbFile string) {
	log.Println("db: connecting")
	_db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println("db: error", err)
		panic(err)
	}

	db = _db
	log.Println("db: connected")

	EnsureSchema(db)
}
