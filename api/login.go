package api

import (
	"fmt"
	"net/http"

	"github.com/facutk/golaburo/utils"
)

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	_, tokenString, _ := utils.TokenAuth.Encode(map[string]interface{}{"user_id": 123})

	fmt.Fprint(w, tokenString)
}
