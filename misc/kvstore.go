package main

import (
    "net"
    "fmt"
    "bufio"
    "strings"
)

var store = make(map[string]string)

func main() {
    li, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }

    for {
        conn, err := li.Accept()
        if err != nil {
            panic(err)
        }

        go handle(conn)
    }
}

func handle(conn net.Conn) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        fs := strings.Fields(ln)
        switch fs[0] {
        case "GET":
            k := fs[1]
            v := store[k]
            fmt.Fprintln(conn, v)
        case "SET":
            k := fs[1]
            v := fs[2]
            store[k] = v
        case "DEL":
            k := fs[1]
            delete(store, k)
        default:
            fmt.Fprintln(conn, "invalid operation")
        }
    }
}
