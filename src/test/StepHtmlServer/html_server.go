package main

import (
    "fmt"
    "net/http"
    "Step"
    "html/template"
    "strconv"
)

type StepShow struct {
    EventTypeId     int
    EventTypeName   string
    EventTypeIcon   string
    StartTime       string
    EndTime         string
}

type WebShow struct {
    EventTypeList []Step.Event
    PsthList []StepShow
}

func (ws *WebShow) LoadStepShow(stepList []Step.Step) {
    for _, v := range stepList {
        var showtem StepShow;
        showtem.EventTypeId = v.TypeId
        event := em.GetEventInfoByType(v.TypeId)
        showtem.EventTypeName = event.TypeName
        showtem.EventTypeIcon = event.TypeIcon
        showtem.StartTime = v.StartTime.String()
        showtem.EndTime = v.EndTime.String()

        ws.PsthList = append(ws.PsthList, showtem)
    }
}

var pl Step.PlayerPsth
var em Step.EventManage

func main() {
    //HtmlServer := http.FileServer(http.Dir("."))
    //http.Handle("/", HtmlServer)
    //http.HandleFunc("/test", testFun)

    var szFileName = "./db.xml";
    pl.LoadFromFile(szFileName)
    em.LoadFromFile();

    http.HandleFunc("/", testFun)
    http.HandleFunc("/BeginEvent", testBegin)

    err := http.ListenAndServe(":8000", nil)
    CheckErr(err)
}

func testBegin(w http.ResponseWriter, r *http.Request) {
    szTypeId := r.URL.Query()["TypeId"][0]
    fmt.Println(r.URL.Query())
    nTypeId,_ := strconv.Atoi(szTypeId)
    pl.AddStep(nTypeId)
    pl.SaveToFile("./db.xml");

    testFun(w, r)
}

func testFun(w http.ResponseWriter, r *http.Request){
    //fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.RawQuery))
    //fmt.Fprintf(w, "hello, %q", r.URL.RawQuery)
    //fmt.Println(r.URL.Query())

    ShowAll(w)
}

func ShowAll(w http.ResponseWriter) {
    v := new (WebShow)
    v.EventTypeList = em.EventList
    v.LoadStepShow(pl.PsthList)

    t, err := template.ParseFiles("./index.html")
    CheckErr(err);

    err = t.Execute(w, v)
    CheckErr(err);
}

/*func ShowAllType(w http.ResponseWriter) {
    for _, v := range em.EventList {
        t := template.New("Event template")
        t, err := t.Parse(temp_a)
        CheckErr(err);

        err = t.Execute(w, v)
        CheckErr(err);
    }
}

func ShowAllList(w http.ResponseWriter) {
    for _, v := range pl.PsthList {
        var showtem StepShow;
        showtem.EventTypeId = v.TypeId
        event := em.GetEventInfoByType(v.TypeId)
        showtem.EventTypeName = event.TypeName
        showtem.EventTypeIcon = event.TypeIcon
        showtem.StartTime = v.StartTime.String()
        showtem.EndTime = v.EndTime.String()

        t := template.New("Psth template")
        t, err := t.Parse(temp)
        CheckErr(err);

        err = t.Execute(w, showtem)
        CheckErr(err);
    }
}*/

func CheckErr(e error) {
    if e != nil {
        fmt.Println("error :", e.Error())
    }
}
