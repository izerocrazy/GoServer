package main

import (
    "fmt"
    "Module"
    "CircleNode"
)

func CreateModule(msg string) Module.Moduler {
    if msg == "CircleModuler" {
        retModule := new (CircleNode.CircleModuler)
        return retModule
    }

    return nil
}

var cm Module.Moduler

func main() {
    // init 
    fmt.Print("main start\n")

    cm = CreateModule("CircleModuler")
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
