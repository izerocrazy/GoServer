package Page

import (
// "fmt"
)

type Thing struct {
	Name      string    `xml:"name"`
	TypeId    int       `xml:"type_id"`
	StartTime time.Time `xml:"start_time"`
	EndTime   time.Time `xml:"end_time"`
}

type Show struct {
	List []Thing `xml:"list"`
}

func (S *Show) Init() {

}
