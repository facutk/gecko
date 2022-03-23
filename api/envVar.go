package api

import (
	"fmt"
	"net/http"

	"github.com/facutk/golaburo/utils"
)

func GetEnv(w http.ResponseWriter, r *http.Request) {
	secretKey := utils.GetEnv("FOO", "")

	fmt.Fprint(w, secretKey)
}
