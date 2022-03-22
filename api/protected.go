package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func GetProtectedHandler(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
}
