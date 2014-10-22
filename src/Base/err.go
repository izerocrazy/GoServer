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

func PrintLog(szErr string) {
	fmt.Println("=======>>> Log Begin <<<=======")
	fmt.Println(szErr)
	fmt.Println("=======>>> Log End <<<=======")
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

func Fmtprintln(i interface{}) {
	fmt.Println("==========>>> Println Beign <<<======")
	fmt.Println(i)
	fmt.Println("==========>>> Println End <<<======")
}
