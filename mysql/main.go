package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

func main() {
    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/lol")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        fmt.Println(err)
    }
}
