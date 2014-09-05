package Todo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"
)

type Log struct {
	LogId    int
	LogDesc  string
	OldEvent Event
}

type EventLog struct {
	EventId int
	LogList []Log
}

type EventLogManager struct {
}
