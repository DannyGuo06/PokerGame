package main

import "fmt"
import "net"

func main() {
    startServer(2)
}

func startServer(numPlayers int) {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("something went wrong:", err)
        return
    }
    fmt.Println("server started")

    var connections []net.Conn
    for i := 0; i < numPlayers; i++ {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("something went wrong:", err)
            return
        }
        connections = append(connections, conn)
        fmt.Println("player", i+1, "connected")
    }
	msg := readMessage(connections[0])
	fmt.Println("player 1 said:", msg)
    for _, conn := range connections {
        conn.Close()
    }
    listener.Close()
}

func readMessage(conn net.Conn) string {
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("something went wrong:", err)
        return "disconnected"
    }
    message := string(buffer[:n])
    return message
}