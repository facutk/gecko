package models

import (
	"log"
)

func GetPragmaSchemaVersion() int {
	rows, err := db.Query("pragma schema_version;")

	if err != nil {
		log.Println("error: ", err.Error())
	}
	var schema_version int
	for rows.Next() { // Iterate and fetch the records from result cursor
		rows.Scan(&schema_version)
		log.Println("schema_version: ", schema_version)
	}

	return schema_version
}
