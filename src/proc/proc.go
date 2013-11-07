package main

import (
    "fmt"
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
    }

    return nil
}

var cm Module.Moduler

func main() {
    // init 
    fmt.Print("main start\n")

    cm = CreateModule("EncodeClient")
    if cm != nil {
        cm.Load()
    }

    // Breath
    Breath()
}

func Breath(){
    if cm.IsSelfRun() == false {
        cm.Breath()
    }
}
