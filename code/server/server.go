package main

import (
    "net"
    "fmt"
    "bufio"
    "strings"
)

func handle(conn net.Conn) {
    defer conn.Close()
    fmt.Println("Client connected")
    for {
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
        break
    }
    fmt.Print("Message Received:", string(message))
    newmessage := strings.ToUpper(message)
    conn.Write([]byte(newmessage + "\n"))
    }
}

func main() {
    fmt.Println("Launching server...")
    
    server, _ := net.Listen("tcp", ":8081")

    for {
        conn, _ := server.Accept()
        if conn != nil {
            go handle(conn)
        }
        conn = nil
    }
}
