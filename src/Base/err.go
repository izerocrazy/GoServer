package Base

import "fmt"

func CheckErr(e error) {
    if e != nil {
        fmt.Println("error :", e.Error())
    }
}
