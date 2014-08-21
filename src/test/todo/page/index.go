package Page

import (
// "encoding/xml"
// "html/template"
// "time"
)

type Index struct {
	TestName string `xml:"testname"`
}

func (I *Index) Init() {
	I.TestName = "Test"
}
