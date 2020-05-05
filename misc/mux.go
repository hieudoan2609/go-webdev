package main

import(
    "fmt"
    "net"
    "bufio"
    "strings"
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

        go request(conn)
    }
}

func request(conn net.Conn) {
    i := 0
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        req := scanner.Text()
        fmt.Println(req)
        if i == 0 {
            mux(conn, req)
        }
        if req == "" {
            break
        }
        i++
    }
}

func mux(conn net.Conn, req string) {
    m := strings.Fields(req)[0]
    u := strings.Fields(req)[1]
    fmt.Println("***METHOD", m)
    fmt.Println("***URI", u)
    switch u {
    case "/":
        index(conn)
    case "/about":
        about(conn)
    case "/contact":
        contact(conn)
    case "/apply":
        switch m {
        case "GET":
            apply(conn)
        case "POST":
            postApply(conn)
        }
    default:
        notFound(conn)
    }
}

func index(conn net.Conn) {
    body := `
        <!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
        <strong>INDEX</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        </body></html>
    `
    respond(conn, body)
}

func about(conn net.Conn) {
    body := `
        <!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
        <strong>ABOUT</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        </body></html>
    `
    respond(conn, body)
}

func contact(conn net.Conn) {
    body := `
        <!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
        <strong>CONTACT</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        </body></html>
    `
    respond(conn, body)
}

func apply(conn net.Conn) {
    body := `
        <!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
        <strong>APPLY</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        <form method="POST" action="/apply">
        <input type="submit" value="apply">
        </form>
        </body></html>
    `
    respond(conn, body)
}

func postApply(conn net.Conn) {
    body := `
        <!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
        <strong>POST APPLY</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        </body></html>
    `
    respond(conn, body)
}

func notFound(conn net.Conn) {
    body := "404 NOT FOUND"
    respond(conn, body)
}

func respond(conn net.Conn, body string) {
    fmt.Fprint(conn, "HTTP/1.1 200 OK\n")
    fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprint(conn, "\n")
    fmt.Fprint(conn, body)
}
