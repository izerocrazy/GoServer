package main

import (
    "fmt"
    "os"
    "Module"
    "CircleNode"
    "EncodeNet"
)

func CreateModule(msg string) Module.Moduler {
    if msg == "CircleModuler" {
        retModule := new (CircleNode.CircleModuler)
        return retModule
    } else if msg == "EncodeClient" {
        retModule := new (EncodeNet.EncodeClient)
        return retModule
    } else if msg == "EncodeServer" {
        retModule := new (EncodeNet.EncodeServer)
        return retModule
    }

    return nil
}

var cm_lst = make([]Module.Moduler, 0)
var c chan int

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s [Module Name]\n", os.Args[0])
        os.Exit(1)
    }

    c = make(chan int)

    // init 
    fmt.Print("main start\n")

    for i := 1; i < len(os.Args); i++ {
        cm := CreateModule(os.Args[i])
        if cm != nil {
            cm_lst = append(cm_lst, cm)
            cm.Load()
        }
    }

    // Breath
    go Breath()

    <-c
}

func Breath(){
    fmt.Print("main breath\n")

    for i := 0; i < len(cm_lst); i++ {
        cm := cm_lst[i]
        if cm.IsSelfRun() == false {
            cm.Breath()
        }
    }

    c <- 1
}
