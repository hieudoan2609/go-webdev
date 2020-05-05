package main

import (
    "net/http"
)

func currentUser(req *http.Request) (string, bool) {
    c, err := req.Cookie("session")
    if err != nil {
        return "", false
    }

    s, ok := sessions[c.Value]
    if !ok {
        return "", false
    }

    _, ok = users[s.Email]
    if !ok {
        return "", false
    }

    return s.Email, true
}