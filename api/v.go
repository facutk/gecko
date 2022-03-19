package api

import (
	"fmt"
	"net/http"

	"github.com/facutk/golaburo/models"
)

func GetMigrationVersion(w http.ResponseWriter, r *http.Request) {
	hits := models.GetMigrationVersion()

	fmt.Fprint(w, hits)
}
