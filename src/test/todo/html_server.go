package main

import (
	"Base"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	// "Step"
	"html/template"
	"test/todo/page"
	// "strconv"
	// "time"
)

func main() {
	myHandler := http.HandlerFunc(Multiplexer)

	err := http.ListenAndServe(":8000", myHandler)
	Base.CheckErr(err)
}

// 强制要求所有都读取 html 文件
func Multiplexer(rw http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	// 如果 url 请求的是 html 则生成对应的 html
	if strings.Contains(request.URL.Path, ".html") {
		fmt.Println("User ask for html", request.URL.Path)

		//打开文件
		szFilePath := "./page" + request.URL.Path
		t, err := template.ParseFiles(szFilePath)
		bSucc := Base.CheckErr(err)
		if bSucc == false {
			rw.WriteHeader(http.StatusNoContent)
		} else {
			v := new(Page.Index)
			v.Init()
			// 按照模板生成 html 文件
			err = t.Execute(rw, v)
			Base.CheckErr(err)
		}
	} else {
		fmt.Println("User ask for server", request.URL.Path)

		//提供服务
		if request.URL.Path == "/add" { // 存入数据
			// 取出数据，存入文件
			// 1 url 中的值
			strMap := request.URL.Query()
			fmt.Println(strMap)
			// 2 其他的值，像是 session（密码）
			// 3 (序列化)转为 xml ，保存
			f := Base.CreateOrAppendFile("db.json")
			defer f.Close()
			encode := json.NewEncoder(f)
			err := encode.Encode(strMap)
			Base.CheckErr(err)
		}
	}
}
