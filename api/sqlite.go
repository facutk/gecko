package api

import (
	"fmt"
	"net/http"

	"github.com/facutk/golaburo/models"
)

func GetSqliteSchemaVersion(w http.ResponseWriter, r *http.Request) {
	schemaVersion := models.GetPragmaSchemaVersion()

	fmt.Fprint(w, schemaVersion)
}
