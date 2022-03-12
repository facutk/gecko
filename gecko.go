package main

import (
    "fmt"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    str := r.URL.Path[1:]
    upperStr := strings.ToUpper(str)
    fmt.Fprintf(w, upperStr)
}
