package main

import "fmt"
import "hash/crc32"
import "net"
import "bufio"

func main() {
    for {
        var s string
        fmt.Print("choose 1 for create Server\n")
        fmt.Print("choose 2 for create player\n")
        fmt.Print("Enter your choose")
        fmt.Scanf("%s", &s)
        fmt.Printf("%s", s)
        if s[0] == '1'{
            fmt.Print("[user]CreateSvr\n")
            go CreateSvr()
        } else {
            fmt.Print("[user]Player\n")
            go CreatePlayer()
        }
    }
}

func CreatePlayer(){
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Print("[player] cant not tcp Dial\n")
        return
    }
    //fmt.Fprintf(conn, "^^^\n")
    nSize, err := conn.Write([]byte("helloKetty\n"))
    if err != nil {
        fmt.Print("[player] can not tcp send\n")
        return
    }

    fmt.Printf("[player]send data size :%d\n", nSize)

    line, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Print("[player] can not tcp read\n")
        return
    }
    fmt.Printf("[player] server send back: %s \n", line)

    defer conn.Close()
}

func CreateSvr(){
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Print("[server]can not set tcp Listen\n")
        return
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Print("[server]can not tcp Accept\n")
            continue
        }

        fmt.Print("[server]Get One Player\n")
        go SerEcho(conn)
    }
}

func SerEcho(conn net.Conn) {
    fmt.Print("server wait for player's Enter")
    //line, err := bufio.NewReader(conn).ReadString('\n')
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

func CalHash(key string) {
    v := crc32.ChecksumIEEE([]byte(key))
    fmt.Println(v)
}

