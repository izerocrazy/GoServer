package Base

import "os"

func CreateOrAppendFile(szFileName string) (file *os.File) {
	//file, err := os.OpenFile(szFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777);
	file, err := os.OpenFile(szFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	CheckErr(err)

	return file
}
