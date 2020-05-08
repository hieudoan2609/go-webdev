package main

import (
    "encoding/json"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "fmt"
    "github.com/hieudoan2609/go-webdev/mvc/models"
)

var users = map[string]models.User{}

func main() {
    r := httprouter.New()
    r.GET("/users/:id", getUser)
    r.POST("/users/:id", createUser)
    r.DELETE("/users/:id", deleteUser)
    http.ListenAndServe(":8080", r)
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    u := users[id]

    uj, _ := json.Marshal(u)
    respond(w, uj)
}

func createUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    u := models.User{}

    json.NewDecoder(r.Body).Decode(&u)
    users[id] = u

    uj, _ := json.Marshal(users[id])
    respond(w, uj)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    u := users[id]

    delete(users, id)

    uj, _ := json.Marshal(u)
    respond(w, uj)
}

func respond(w http.ResponseWriter, body []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%s\n", body)
}
