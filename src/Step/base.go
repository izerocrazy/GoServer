package Step

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"time"
)

const temp = `TypeId={{.TypeId}}, StartTime={{.StartTime}}, EndTime = {{.EndTime}}`

type Step struct {
	TypeId    int       `xml:"type_id"`
	StartTime time.Time `xml:"start_time"`
	EndTime   time.Time `xml:"end_time"`
}

type PlayerPsth struct {
	PsthList []Step `xml:"psth_list"`
}

func (p *PlayerPsth) AddStep(nTypeId int) {
	i := len(p.PsthList)
	now := time.Now()
	newPsth := Step{
		TypeId:    nTypeId,
		StartTime: now,
		EndTime:   now,
	}
	if i > 0 {
		p.PsthList[i-1].EndTime = now
	}

	p.PsthList = append(p.PsthList, newPsth)
}

func (p *PlayerPsth) LoadFromFile(szFileName string) {
	buf, err := ioutil.ReadFile(szFileName)
	if len(buf) == 0 {
		return
	}

	CheckErr(err)

	err = xml.Unmarshal(buf, &p)

	CheckErr(err)
}

func (p *PlayerPsth) SaveToFile(szFileName string) {
	buf, err := xml.Marshal(p)
	CheckErr(err)

	ioutil.WriteFile(szFileName, buf, 0777)
}

func (p *PlayerPsth) ShowTemplate() {
	for i := 0; i < len(p.PsthList); i++ {
		psth := p.PsthList[i]

		t := template.New("Psth template")
		t, err := t.Parse(temp)
		CheckErr(err)

		err = t.Execute(os.Stdout, psth)
		CheckErr(err)
	}
}

func PsthTest() {
	var MyPsthList PlayerPsth
	var szFileName = "./db.xml"

	MyPsthList.LoadFromFile(szFileName)
	fmt.Println(len(MyPsthList.PsthList))

	if len(MyPsthList.PsthList) == 0 {
		MyPsthList.AddStep(1)
	}

	MyPsthList.ShowTemplate()
	MyPsthList.SaveToFile(szFileName)
}
