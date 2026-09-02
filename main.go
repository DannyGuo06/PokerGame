package main

import "fmt"
import "net"
import "strings"

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
	gameLoop(connections) // calls connections to keep reading and sending msg
    for _, conn := range connections {
        conn.Close()
    }
    listener.Close()
}

func readMessage(conn net.Conn) string { //recieve msg
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("something went wrong:", err)
        return "disconnected"
    }
    message := strings.TrimSpace(string(buffer[:n]))
    return message
}

func sendMessage(conn net.Conn, message string) { //send msg
    _, err := conn.Write([]byte(message))
    if err != nil {
        fmt.Println("something went wrong:", err)
    }
}

func gameLoop(connections []net.Conn) { // recieve and send msg
    i := 0
    for {
        current := i % 2
        other := (i + 1) % 2

        msg := readMessage(connections[current])
        fmt.Println("player", current+1, "said:", msg)

        if msg == "fold" || msg == "disconnected" {
            sendMessage(connections[other], "opponent folded")
            break
        }

        sendMessage(connections[other], msg)
        i++
    }
}