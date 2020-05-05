package main

import (
    "github.com/satori/go.uuid"
    "golang.org/x/crypto/bcrypt"
    "html/template"
    "net/http"
    "time"
    "fmt"
)

type session struct {
    Email string
    Expiry time.Time
}

var users = map[string][]byte{} // email: password
var sessions = map[string]session{} // sid: session
var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("view/*"))
}

func main() {
    go cleanSessionsRoutine()

    http.HandleFunc("/", index)
    http.HandleFunc("/signup", signup)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)
    http.ListenAndServe(":8080", nil)
}

func cleanSessionsRoutine() {
    for {
        for k, v := range sessions {
            if time.Now().After(v.Expiry) {
                delete(sessions, k)
                fmt.Println("cleaned up " + v.Email)
            }
        }

        time.Sleep(time.Second / 2)
    }
}

func index(w http.ResponseWriter, req *http.Request) {
    u, _ := currentUser(req)
    tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
    var message string

    if req.Method == http.MethodPost {
        email := req.FormValue("email")
        pwd := req.FormValue("password")
        _, ok := users[email]
        if ok {
            message = "email has been taken"
        } else {
            digest, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
            if err != nil {
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
            users[email] = digest
            message = "user successfully registered"
        }
    }

    tpl.ExecuteTemplate(w, "signup.gohtml", message)
}

func login(w http.ResponseWriter, req *http.Request) {
    var message string

    if req.Method == http.MethodPost {
        email := req.FormValue("email")
        pwd := req.FormValue("password")
        u, ok := users[email]
        if !ok {
            message = "user does not exist"
        } else {
            err := bcrypt.CompareHashAndPassword(u, []byte(pwd))
            if err != nil {
                message = "password is incorrect"
            } else {
                sid, _ := uuid.NewV4()
                c := &http.Cookie{
                    Name: "session",
                    Value: sid.String(),
                }
                s := session{
                    Email: email,
                    Expiry: time.Now().Add(time.Second * 10),
                }
                http.SetCookie(w, c)
                sessions[c.Value] = s
                message = "logged in successfully"
            }
        }
    }

    tpl.ExecuteTemplate(w, "login.gohtml", message)
}

func logout(w http.ResponseWriter, req *http.Request) {
    var message string

    if req.Method == http.MethodPost {
        u, ok := currentUser(req)
        if !ok {
            message = "not logged in"
        } else {
            c := &http.Cookie{
                Name: "session",
                Value: "",
                MaxAge: -1,
            }
            delete(sessions, u)
            http.SetCookie(w, c)
            message = "logged out successfully"
        }
    }

    tpl.ExecuteTemplate(w, "logout.gohtml", message)
}
