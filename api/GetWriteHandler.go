package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetWriteHandler(w http.ResponseWriter, r *http.Request) {
	str := chi.URLParam(r, "str")
	fmt.Fprint(w, str)

	d1 := []byte(str)
	err := os.WriteFile("./dat", d1, 0644)
	check(err)
}
