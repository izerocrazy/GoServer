package EncodeNet

import (
    "fmt"
    "net"
    "os"
    "encoding/json"
)

type Server interface {
    Init()
    Run()
}

// ========
type EncodeServer struct {
    // For Link Client
    ServerListenInfo string
    ServerTCPAddr *net.TCPAddr
    ServerListen *net.TCPListener
}

func (es *EncodeServer) Init() {
    var str string = "127.0.0.1:8282"
    es.setServerListenString(str)
    es.createServerListen()
}

func (es *EncodeServer) Breath() {
    fmt.Printf("EncodeServer: Breath\n")
}

func (es *EncodeServer) Run() {
    fmt.Printf("EncodeServer: Run\n")
    go es.runFunc()
}

func (es *EncodeServer) runFunc() {
    for {
        conn, err := es.ServerListen.Accept()
        if err != nil {
            fmt.Fprintf(os. Stderr, "can not Accept err: %s\n", err.Error())
        }

        if conn != nil{
            go doConn(conn)
        }
    }
}

func (es *EncodeServer) Stop() {
    fmt.Printf("EncodeServer:Stop\n")

    es.ServerListen.Close()
}

func (es *EncodeServer) IsSelfRun() bool {
    fmt.Printf("EncodeServer:IsSelfRun\n")

    return true
}

func (es *EncodeServer) Load() error {
    fmt.Printf("EnocodeServer:Load\n")
    es.Init()

    if es.IsSelfRun() == true {
        es.Run()
    }

    return nil
}

func (es *EncodeServer) Unload() error {
    fmt.Printf("EncodeServer:Unload\n")

    if es.IsSelfRun() == true {
        es.Stop()
    }

    return nil
}

func (es *EncodeServer) setServerListenString(str string) {
    es.ServerListenInfo = str
}

func (es *EncodeServer) createServerListen() {
    var err error
    es.ServerTCPAddr, err = net.ResolveTCPAddr("tcp4", es.ServerListenInfo)
    if err != nil {
        fmt.Fprintf(os.Stderr, "can not ResolveTCPAddr err: %s\n", err.Error())
        os.Exit(1)
    }

    es.ServerListen, err = net.ListenTCP("tcp", es.ServerTCPAddr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "can not ListenTCP err :%s\n", err.Error())
        os.Exit(1)
    }
}

func doConn(conn net.Conn) {
    fmt.Printf("Get one conn\n")

    //var buf [1024]byte
    /*_, err := conn.Read(buf[0:])
    if err!= nil {
        return
    }*/

    // 解码TCPAddr
    //var tcpAddr net.TCPAddr
    var testData string
    testData = "ip: 127.0.0.1:8282"
    decoder := json.NewDecoder(conn)
    decoder.Decode(&testData)
    
    defer conn.Close()
}

/*func main(){
    if len(os.Args) < 1 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port host:port\n", os.Args[0])
        os.Exit(1)
    }

    if len(os.Args) == 3 {
        link := os.Args[2]
        myTCPString := os.Args[1]
        TellLinkImCome(link, myTCPString)
    }

    service := os.Args[1]

    listener := CreateServerListen(service)

    if listener != nil {
        DoServer(listener)
    }
}*/

