package main

import (
    "net/http"
    "io"
    "io/ioutil"
    "fmt"
)

func main() {
    // http.Handle
    http.HandleFunc("/", serveForm)
    http.ListenAndServe(":8080", nil)
}

func serveForm(w http.ResponseWriter, req *http.Request) {
    var form = `
    <form method="POST" enctype="multipart/form-data">
    <input type="file" name="upload">
    <input type="submit">
    </form>
    <br>
    `

    if req.Method == http.MethodPost {
        f, _, _ := req.FormFile("upload")
        defer f.Close()
        bs, _ := ioutil.ReadAll(f)
        fmt.Println(string(bs))
        form += string(bs)
    }

    w.Header().Set("Content-Type", "text/html")
    io.WriteString(w, form)
}