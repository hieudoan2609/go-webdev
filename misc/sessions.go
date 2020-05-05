package main

import (
    "github.com/satori/go.uuid"
    "html/template"
    "net/http"
)

type user struct {
    UserName string
    Email string
}

var tpl *template.Template
var users = map[string]user{}

func init() {
    tpl = template.Must(template.ParseFiles("sessions.gohtml"))
}

func main() {
    http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
    c, err := req.Cookie("session")
    if err != nil {
        id, _ := uuid.NewV4()
        c = &http.Cookie{
            Name: "session", 
            Value: id.String(),
        }
        http.SetCookie(w, c)
    }

    if req.Method == http.MethodPost {
        users[c.Value] = user{
            UserName: req.FormValue("username"),
            Email: req.FormValue("email"),
        }
    }

    u := users[c.Value]
    tpl.ExecuteTemplate(w, "sessions.gohtml", u)
}
