package Base

import (
	"html/template"
	"os"
)

func CreateOrAppendFile(szFileName string) *os.File {
	//file, err := os.OpenFile(szFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777);
	file, err := os.OpenFile(szFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	CheckErrExit(err)

	return file
}

func CreateOrAppendFileWithTemplate(szFileName string, szTemplateFileName string, data interface{}) *os.File {
	file := CreateOrAppendFile(szFileName)
	t, err := template.ParseFiles(szTemplateFileName)
	CheckErrExit(err)

	err = t.Execute(file, data)
	CheckErrExit(err)

	return file
}
