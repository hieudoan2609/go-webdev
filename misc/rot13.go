package main

import (
    "bufio"
    "net"
    "fmt"
)

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
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        bs := []byte(ln)
        rot13(bs)
        fmt.Fprintf(conn, "%s\n", bs)
    }
}

func rot13(bs []byte) {
    for i, v := range bs {
        if v <= 109 {
            bs[i] = v + 13
        } else {
            bs[i] = v - 13
        }
    }
}
