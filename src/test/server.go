package main 

import "fmt"
//import "hash/crc32"
import "net"

func main() {
    ln := SetServer()
    for ln != nil {
        fmt.Printf("server is waiting for player\n")
        conn, err := ln.Accept()
        if err != nil {
            fmt.Printf("serve can not get this player\n")
            continue
        }

        go SerEcho(conn)
    }
}

func SetServer() net.Listener{
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Printf("server can not set this listen on the port 8080\n")
        return nil
    }

    return ln
}

func SerEcho(conn net.Conn) {
    fmt.Printf("Server get one player and waiting for his' enter\n")
    for {
        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Printf("[server]Failuer to reader: %s\n", err.Error())
            return
        }
        fmt.Printf("[server]Get Player Send Data: %s, size: %d\n", string(buf), n)

        _, err = conn.Write(buf)
        if err != nil {
            fmt.Printf("[server]Failuer to write: %s\n", err.Error())
            return
        }
    }
}
