package EncodeServer

import (
    "fmt"
    "net"
    "os"
    "encoding/json"
)

type MyData struct {
    Ip string
    Port string
}

type EncodeServer struct {
    // For Link Client
    ServerListenInfo string
    ServerTCPAddr *net.TCPAddr
    ServerListen *net.TCPListener
}

func (es *EncodeServer) InitServer(str string) {
    es.setServerListenString(str)
    es.createServerListen()
}

func (es *EncodeServer) RunServer() {
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
    var testData MyData
    decoder := json.NewDecoder(conn)
    decoder.Decode(&testData)
    fmt.Printf("ip: %s, port, %s", testData.Ip, testData.Port)

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

type EncodeClient struct {
    ClientTCPInfo string
    ClientConn net.Conn
}

func (ec *EncodeClient) InitClient(server string) {
    ec.ClientTCPInfo = server
}

func (ec *EncodeClient) ConnectServer() {
    var err error
    ec.ClientConn, err = net.Dial("tcp", ec.ClientTCPInfo)
    if err != nil {
        fmt.Fprintf(os.Stderr, "can not DialTCP err: %s \n", err.Error)
        os.Exit(1)
    }
}

func (ec *EncodeClient) doClient() {
    var testData MyData
    testData.Ip = ec.ClientTCPInfo
    testData.Port = ec.ClientTCPInfo

    fmt.Printf("to send my encoder\n")
    encoder := json.NewEncoder(ec.ClientConn)
    encoder.Encode(&testData)
    fmt.Printf("send my encoder over %s, %s\n", testData.Ip, testData.Port)
}
