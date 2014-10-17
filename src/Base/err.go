package Base

import (
	"fmt"
	"os"
)

func CheckErr(e error) bool {
	if e != nil {
		fmt.Println("error :", e.Error())
		return false
	}

	return true
}

func CheckErrExit(e error) bool {
	if CheckErr(e) == false {
		os.Exit(0)
		return false
	}

	return true
}

func PrintErr(szErr string) {
	fmt.Println("=======>>> Error Begin <<<=======")
	fmt.Println(szErr)
	fmt.Println("=======>>> Error End <<<=======")
}

func PrintErrExit(szErr string) {
	PrintErr(szErr)

	os.Exit(0)
}
