package api

import (
	"fmt"
	"os"

	"net/http"
)

func GetReadHandler(w http.ResponseWriter, r *http.Request) {
	dat, err := os.ReadFile("./dat")
	check(err)
	fmt.Print(string(dat))

	fmt.Fprint(w, string(dat))
}
