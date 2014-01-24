package main

import (
    "fmt"
    "time"
    "os"
    "bufio"
    "html/template"
    "encoding/xml"
)

//const temp = `TypeId={{.TypeId}}, StartTime={{.StartTime}}, EndTime = {{.EndTime}}`

type Psth struct {
    TypeId      int         `xml:"type_id"`
    StartTime   time.Time   `xml:"start_time"`
    EndTime     time.Time   `xml:"end_time"`
}

func (p Psth) WriteFile() {
    f, _ := os.OpenFile("./test.db", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
    defer f.Close()

    /*t := template.New("Psth template 2")
    t, err := t.Parse(temp)
    CheckErr(err)

    err = t.Execute(f, psth)
    CheckErr(err)*/

    w := bufio.NewWriter(f)     // NewWirter() 和 io.Writer，以及 os.Stdout 的关系同上
    defer w.Flush()

    buf,err := xml.Marshal(p)
    CheckErr(err)

    w.Write(buf)
}

type PlayerPsth struct {
    PsthList    []Psth      `xml:"psth_list"`
}

func (p PlayerPsth) AddStep(nTypeId int){
    i := len(p.PsthList)
    now := time.Now()
    newPsth := Psth{
        TypeId:     nTypeId,
        StartTime:  now,
        EndTime:    now,
    }
    if i > 0 {
        p.PsthList[i-1].EndTime = now
    }

    p.PsthList = append(p.PsthList, newPsth)
}


func ReadFile(MyPsthList []Psth) ([]Psth){
    f, _ := os.OpenFile("./test.db", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
    defer f.Close()

    buf := make([]byte, 1024)
    n, err := f.Read(buf)
    CheckErr(err)

    err = xml.Unmarshal(buf)
}
func CheckErr(err error){
    if err != nil {
        fmt.Println(err);
    }
}

func ShowTemplate(psth Psth) () {
    t := template.New("Psth template")
    t, err := t.Parse(temp)
    CheckErr(err);

    err = t.Execute(os.Stdout, psth)
    CheckErr(err);
}

func main() {
    var MyPsthList PlayerPsth
}

