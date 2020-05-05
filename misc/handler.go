package main

import (
    "io"
    "net/http"
)

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
    case "/dog":
        io.WriteString(w, "dog dog dog")
    case "/cat":
        io.WriteString(w, "cat cat cat")
    }
}

func main() {
    var h handler
    http.ListenAndServe(":8080", h)
}