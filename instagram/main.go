package main

import (
    "crypto/sha1"
    "github.com/satori/go.uuid"
    "net/http"
    "html/template"
    "fmt"
    "strings"
    "io"
    "os"
    "path/filepath"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("views/*"))
}

func main() {
    http.HandleFunc("/", index)
    http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
    http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
    c := getCookie(req)

    if req.Method == http.MethodPost {
        mf, fh, err := req.FormFile("photo")
        check(err)
        defer mf.Close()

        ext := strings.Split(fh.Filename, ".")[1]
        h := sha1.New()
        io.Copy(h, mf)
        fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

        wd, err := os.Getwd()
        check(err)
        path := filepath.Join(wd, "public", fname)
        nf, err := os.Create(path)
        check(err)
        defer nf.Close()

        mf.Seek(0, 0)
        io.Copy(nf, mf)

        c = appendPhoto(w, c, fname)
    }

    photos := strings.Split(c.Value, ",")[1:]
    fmt.Println(photos)
    tpl.ExecuteTemplate(w, "index.gohtml", photos)
}

func getCookie(req *http.Request) *http.Cookie {
    c, err := req.Cookie("session")
    if err != nil {
        sid, _ := uuid.NewV4()
        c = &http.Cookie{
            Name: "session",
            Value: sid.String(),
        }
    }
    return c
}

func appendPhoto(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
    val := c.Value
    val += "," + fname
    c.Value = val
    http.SetCookie(w, c)
    return c
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
