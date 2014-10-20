package main

import (
	"fmt"
	"strings"
)

func main() {
	tbMain := strings.Split("A", ":")
	fmt.Println(len(tbMain), tbMain)
	tbMain = strings.Split(":A", ":")
	fmt.Println(len(tbMain), tbMain)
	tbMain = strings.Split("A:", ":")
	fmt.Println(len(tbMain), tbMain)
	tbMain = strings.Split(":", ":")
	fmt.Println(len(tbMain), tbMain)
}
