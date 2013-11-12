package EncodeNet

import (
    "fmt"
    "net"
    "os"
    "encoding/json"
)

type Client interface {
    Init()
    SendData(v interface{}) error
}

//==========
type EncodeClient struct {
    ClientTCPInfo string
    ClientConn net.Conn
    ClientError error
}

func (ec *EncodeClient) Init() {
    fmt.Printf("EncodeClient: Init()")

    ec.ClientTCPInfo = "127.0.0.1:8282"
    ec.connectServer()
}

func (ec *EncodeClient) Breath() {
    fmt.Printf("EncodeClient:Breath()")
}

func (ec *EncodeClient) Run() {
    fmt.Printf("EncodeClient:Run()")

    ec.SendData("EnocodeClient:Run()")
}

func (ec *EncodeClient) Stop() {
    fmt.Printf("EncodeClient:Stop()")

    ec.ClientConn.Close()
}

func (ec *EncodeClient) IsSelfRun() bool {
    return true
}

func (ec *EncodeClient) Load() error {
    fmt.Printf("EncodeClient:Load()")

    if ec.ClientError != nil {
        return ec.ClientError
    }

    ec.Init()
    if ec.IsSelfRun() == true {
        ec.Run()
    }

    return nil
}

func (ec *EncodeClient) Unload() error {
    fmt.Printf("EncodeClient:Unload()")
    if ec.ClientError != nil {
        return ec.ClientError
    }

    if ec.IsSelfRun() == true {
        ec.Stop()
    }

    return nil
}

func (ec *EncodeClient) connectServer() error {
    var err error
    ec.ClientConn, err = net.Dial("tcp", ec.ClientTCPInfo)
    if err != nil {
        fmt.Fprintf(os.Stderr, "can not DialTCP err: %s \n", err.Error())
        ec.ClientError = err
    }

    return err
}

func (ec *EncodeClient) SendData(v interface{}) error {
    if ec.ClientError != nil {
        return ec.ClientError
    }

    fmt.Printf("to send my encoder\n")
    encoder := json.NewEncoder(ec.ClientConn)
    encoder.Encode(&v)
    fmt.Printf("send my encoder over \n")

    return nil
}
