package Base

import (
	"fmt"
)

func CheckErr(e error) bool {
	if e != nil {
		fmt.Println("error :", e.Error())
		return false
	}

	return true
}
