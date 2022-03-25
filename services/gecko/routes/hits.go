package routes

import (
	"fmt"
	"net/http"

	"github.com/facutk/golaburo/models"
)

func GetHits(w http.ResponseWriter, r *http.Request) {
	hits := models.GetHits()

	fmt.Fprint(w, hits)
}
