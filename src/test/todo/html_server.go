package main

import (
	"Base"
	"fmt"
	"net/http"
	"strings"
	// "Step"
	"html/template"
	// "strconv"
	// "time"
)

func main() {
	myHandler := http.HandlerFunc(Multiplexer)

	err := http.ListenAndServe(":8000", myHandler)
	Base.CheckErr(err)
}

func Multiplexer(rw http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	// 如果 url 请求的是 html 则生成对应的 html
	if strings.Contains(request.URL.Path, ".html") {
		fmt.Println("User ask for html", request.URL.Path)

		//打开文件
		szFilePath := "." + request.URL.Path
		t, err := template.ParseFiles(szFilePath)
		// _, err := template.ParseFiles(szFilePath)
		bSucc := Base.CheckErr(err)
		if bSucc == false {
			rw.WriteHeader(http.StatusNoContent)
		} else {
			// 按照模板生成 html 文件
			err = t.Execute(rw, nil)
			Base.CheckErr(err)
		}
	} else {
		fmt.Println("User ask for server", request.URL.Path)

		//提供服务
		if request.URL.Path == "/add" { // 存入数据
			// 取出数据，存入文件
		}
	}
}
