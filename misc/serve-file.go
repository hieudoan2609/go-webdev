package main

import (
    "os"
    "io"
    "net/http"
)

func main() {
    http.HandleFunc("/", girl)
    http.HandleFunc("/girl.webp", girlPic)
    http.ListenAndServe(":8080", nil)
}

func girl(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    io.WriteString(w, `
    <img src="/girl.webp" />
    `)
}

func girlPic(w http.ResponseWriter, req *http.Request) {
    f, err := os.Open("ariel.webp")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    io.Copy(w, f)
}
