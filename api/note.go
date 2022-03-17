package api

import (
	"encoding/json"
	"fmt"
	"os"

	"net/http"
)

func check(e error, w http.ResponseWriter) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		panic(e)
	}
}

type Note struct {
	Message string `json:"message"`
}

func NoteGetHandler(w http.ResponseWriter, r *http.Request) {
	message := ""
	dat, _ := os.ReadFile("/data/dat")
	message = string(dat)

	n := Note{}
	n.Message = message
	fmt.Print(n.Message)

	marshalledTodo, _ := json.MarshalIndent(n, "", "  ")

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(marshalledTodo))
}

func NotePostHandler(w http.ResponseWriter, r *http.Request) {
	var n Note

	err := json.NewDecoder(r.Body).Decode(&n)
	check(err, w)

	d1 := []byte(n.Message)
	errWrite := os.WriteFile("/data/dat", d1, 0644)
	check(errWrite, w)

	w.WriteHeader(200)
}
