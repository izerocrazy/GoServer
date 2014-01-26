package Step

import (
    "fmt"
    "encoding/xml"
    "io/ioutil"
)

type Event struct {
    TypeId      int     `xml:"type_id"`
    TypeName    string  `xml:"type_name"`
    TypeIcon    string  `xml:"type_icon"`
}

type EventManage struct {
    EventList   []Event `xml:"event_list"`
}

func (e *EventManage)LoadFromFile(){
    buf, err := ioutil.ReadFile("./event.xml")
    if len(buf) == 0 {
        return
    }

    CheckErr(err)

    err = xml.Unmarshal(buf, &e)

    CheckErr(err)
}

func (e *EventManage) GetEventInfoByType(nTypId int) (*Event){
    for _, v := range e.EventList {
        if nTypId == v.TypeId {
            return &v
        }
    }

    return nil
}

func CheckErr(err error) {
    if err != nil {
        fmt.Println("err: ", err);
    }
}

func EventTest(){
    var em EventManage;
    em.LoadFromFile();

    fmt.Println(em.EventList);
}
