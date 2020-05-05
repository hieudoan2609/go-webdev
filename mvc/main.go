package main

import (
    "encoding/json"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "fmt"
)

var users = map[string]models.User{}

func main() {
    r := httprouter.New()
    r.GET("/users/:id", getUser)
    r.POST("/users", createUser)
    r.DELETE("/users/:id", deleteUser)
    http.ListenAndServe(":8080", r)
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    u := models.User{
        Name: "hieu",
        Email: "hieudoan2609@gmail.com",
    }

    uj, err := json.Marshal(u)
    check(err)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOk)
    fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
