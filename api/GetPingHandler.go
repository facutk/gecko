package api

import (
	"fmt"
	"net/http"
)

func GetPingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
