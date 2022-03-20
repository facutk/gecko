package api

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetUppercaseHandler(w http.ResponseWriter, r *http.Request) {
	str := chi.URLParam(r, "str")
	upperStr := strings.ToUpper(str)
	w.Write([]byte(upperStr))
}
