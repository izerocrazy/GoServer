package main

import (
	"Step"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type StepShow struct {
	EventTypeId   int
	EventTypeName string
	EventTypeIcon string
	StartTime     string
	EndTime       string
	LastTime      string
}

type WebShow struct {
	EventTypeList []Step.Event
	PsthList      []StepShow
	NowStep       StepShow
}

func (ws *WebShow) LoadOneStepShow(s Step.Step) StepShow {
	var showtem StepShow

	showtem.EventTypeId = s.TypeId
	event := em.GetEventInfoByType(s.TypeId)
	showtem.EventTypeName = event.TypeName
	showtem.EventTypeIcon = event.TypeIcon
	const layout = time.RFC850
	showtem.StartTime = s.StartTime.Format(layout)
	if s.StartTime == s.EndTime {
		s.EndTime = time.Now()
	}
	showtem.EndTime = s.EndTime.Format(layout)
	showtem.LastTime = ws.GetTimeString(s.EndTime.Sub(s.StartTime))

	return showtem
}

func (ws *WebShow) GetTimeString(t time.Duration) string {
	nHour := int(t.Hours())
	nMinute := t.Minutes()
	s := fmt.Sprintf("%d : %d", nHour, int(nMinute)-nHour*60)

	return s
}

func (ws *WebShow) LoadStepShow(stepList []Step.Step) {
	nLen := len(stepList)
	nTimeNow := time.Now()
	for k, _ := range stepList {
		v := stepList[nLen-k-1]
		// 第一个也就是当前的事件
		if k == 0 {
			ws.NowStep = ws.LoadOneStepShow(v)

			timeSince := time.Since(v.StartTime)
			ws.NowStep.LastTime = ws.GetTimeString(timeSince)

			continue
		}

		if nTimeNow.Sub(v.EndTime).Hours() > 24 {
			continue
		}

		ws.PsthList = append(ws.PsthList, ws.LoadOneStepShow(v))
	}
}

var pl Step.PlayerPsth
var em Step.EventManage

func main() {
	HtmlServer := http.FileServer(http.Dir("."))
	http.Handle("/", HtmlServer)
	//http.HandleFunc("/test", testFun)

	var szFileName = "./db.xml"
	pl.LoadFromFile(szFileName)
	em.LoadFromFile()

	// fmt.Println(pl)
	// fmt.Println(em)

	http.HandleFunc("/index", Index)
	http.HandleFunc("/BeginEvent", testBegin)

	err := http.ListenAndServe(":8000", nil)
	CheckErr(err)
}

func testBegin(w http.ResponseWriter, r *http.Request) {
	szTypeId := r.URL.Query()["TypeId"][0]
	fmt.Println(r.URL.Query())
	nTypeId, _ := strconv.Atoi(szTypeId)
	pl.AddStep(nTypeId)
	pl.SaveToFile("./db.xml")

	testFun(w, r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	ShowAll(w)
}

func testFun(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.RawQuery))
	//fmt.Fprintf(w, "hello, %q", r.URL.RawQuery)
	//fmt.Println(r.URL.Query())

	ShowAll(w)
}

func ShowAll(w http.ResponseWriter) {
	v := new(WebShow)
	v.EventTypeList = em.EventList
	v.LoadStepShow(pl.PsthList)

	t, err := template.ParseFiles("./index.html")
	CheckErr(err)

	err = t.Execute(w, v)
	CheckErr(err)
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
