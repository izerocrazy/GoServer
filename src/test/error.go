package main

import "fmt"
import "errors"

func main() {
    var e error;
    e = errors.New("i'm error")

    fmt.Printf("error : %s\n", e.Error())
}
