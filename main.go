package main

import (
    "fmt"
    "net/http"
)

func main() {

    http.Handle("/", http.FileServer(http.Dir("./s/")))
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}