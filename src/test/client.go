package main

import "fmt"
//import "hash/crc32"
import "net"

func main() {
    conn := connectToServer()
    if conn != nil {
        for {
            fmt.Print("Enter your name\n")
            var name string
            fmt.Scanf("%s", &name)

            go SendData(conn, "login")
            go ReadData(conn)
        }
    }
    defer conn.Close()
}

func connectToServer() net.Conn {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Print("Can not connnet to server: localhost : 8080\n")
        return nil
    }
    return conn
}

func SendData(conn net.Conn, msg string) {
    nSize, err := conn.Write([]byte(msg))
    if err != nil {
        fmt.Print("Can not send data to server\n")
    }
    fmt.Printf("[player]send data %s, size :%d\n", msg, nSize)
}

func ReadData(conn net.Conn) {
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Print("can not tcp read\n")
        return 
    }
    fmt.Printf("server send back: %s , size: %d\n", string(buf), n)
}
