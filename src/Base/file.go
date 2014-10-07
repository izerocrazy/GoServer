package Base

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"os"
)

// 此函数用于得到一个文件的句柄，如果文件存在，游标就在最后，如果文件不存在，那么就创建
func CreateOrAppendFile(szFileName string) *os.File {
	//file, err := os.OpenFile(szFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777);
	file, err := os.OpenFile(szFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	CheckErrExit(err)

	return file
}

// 此函数用于得到一个文件的句柄，如果文件存在，游标则在最后，并且按照模板文件在后面追加，如果文件不存在那么就按照模板文件创建
func CreateOrAppendFileWithTemplate(szFileName string, szTemplateFileName string, data interface{}) *os.File {
	file := CreateOrAppendFile(szFileName)
	t, err := template.ParseFiles(szTemplateFileName)
	CheckErrExit(err)

	err = t.Execute(file, data)
	CheckErrExit(err)

	return file
}

func SaveStructToXML(szFileName string, p interface{}) {
	buf, err := xml.Marshal(p)
	CheckErr(err)

	ioutil.WriteFile(szFileName, buf, 0777)
}
