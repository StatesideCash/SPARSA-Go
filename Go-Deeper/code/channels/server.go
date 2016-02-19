package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
    "time"
)

var connStack = make(chan net.Conn)
var isReady = make(chan bool)

func receiveConn(conn net.Conn) {
	fmt.Println("Client connected")
	for {
		connStack <- conn
	}
}

func handleConns() {
    fmt.Println("Simulating processor-intensive task")
    time.Sleep(5 * time.Second)
    fmt.Println("Intensive task done simulating")
    isReady <- true // Telling parent process we are ready

	for {
		conn := <-connStack
        message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Close()
            continue
		}
        fmt.Print("Message Received:", string(message))
        newmessage := strings.ToUpper(message)
        conn.Write([]byte(newmessage + "\n"))
        conn.Close()
	}
}

func main() {
    go handleConns() // Launch the handler
    <-isReady        // Wait until the initial task is finished

	fmt.Println("Launching server...")
	server, _ := net.Listen("tcp", ":8081")
	for {
		conn, _ := server.Accept()
		if conn != nil {
			go receiveConn(conn)
		}
		conn = nil
	}
}
