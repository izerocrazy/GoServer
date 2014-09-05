package Todo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"
)

const (
	E_ING = itoa
	E_Finish
	E_Delete
)

type Event struct {
	EventId    int
	Name       string
	Desc       string
	AddTime    time.Time
	RollTime   time.Time
	DeleteTime time.Time
	FinishTime time.Time
	Vesion     int

	State int
}

func (e *Evnet) SetData(EventId int, Name string, Desc string) {
	e.EventId = EventId
	e.Name = Name
	e.Desc = Desc

	e.AddTime = time.Now()
	e.Vesion = 0

	e.State = E_ING
}

func (e *Event) ChangeRollTime(NewRollTime time.Time) {
	e.RollTime = NewRollTime
}

func (e *Event) FinishEvent() {
	e.FinishTime = time.Now()

	e.State = E_Finish
}

type EventManager struct {
	EventList []*Event
}

func (em *EventManage) FindEvent(nEventId int) *Event {
	for _, e := range em.Eventlist {
		if e.EventId == nEventId {
			return e
		}
	}

	return nil
}
