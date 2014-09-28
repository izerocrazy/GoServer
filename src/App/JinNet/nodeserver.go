package main

import (
    "fmt"
    "CircleNode"
    "os"
    "EncodeServer"
    "hash/crc32"
)

type NodeServer struct {
    cnc CircleNode.NodeCircle
    ees EncodeServer.EncodeServer
    eec EncodeServer.EncodeClient
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port [host:port]\n", os.Args[0])
        os.Exit(1)
    }

    var ns NodeServer;
    ns.ees.InitServer(os.Args[1])
    go ns.ees.RunServer()

    switch {
    case len(os.Args) == 2:
        ns.cnc.InitCircle()
        nHashValue := crc32.ChecksumIEEE([]byte(os.Args[1]))
        ns.cnc.AddNode(int(nHashValue), os.Args[1])
    case len(os.Args) == 3:
        //linkName := os.Args[2]
        //myName := os.Args[1]
        ns.eec.InitClient(os.Args[2])
        ns.eec.ConnectServer()
    }
}
