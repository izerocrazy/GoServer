package main

import (
	"Base"
	"fmt"
)

type Pagename struct {
	Name string
}

func main() {
	fmt.Println("输入名称：")

	var szStr string
	fmt.Scanf("%s", &szStr)

	// .go
	var szFileName = "../page/" + szStr + ".go"
	var szTemplateFileName = "../page/template/template.go"
	data := new(Pagename)
	data.Name = szStr
	Base.CreateOrAppendFileWithTemplate(szFileName, szTemplateFileName, data)

	// .html
	szFileName = "../page/" + szStr + ".html"
	szTemplateFileName = "../page/template/template.html"
	Base.CreateOrAppendFileWithTemplate(szFileName, szTemplateFileName, data)
}
