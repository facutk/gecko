package models

import "log"

func GetHits() int {
	row := db.QueryRow("select hits from visits where id=1;")

	var hits int
	row.Scan(&hits)
	log.Println("hits: ", hits)

	_, err := db.Exec("update visits set hits=$1 where id=1", hits+1)
	if err != nil {
		log.Println("hits: update failed", err)
	}

	return hits
}
