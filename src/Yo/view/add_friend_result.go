package view

import (
	"encoding/json"
	"net/http"
)

type AddFriendData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type IAddFriend struct {
	Retcode  int           `json:"retcode"`
	Msg      string        `json:"msg"`
	Datetime int           `json:"datetime"`
	Data     AddFriendData `json:"data"`
}

type AddFriendResult struct {
}

func (fr *AddFriendResult) Render(i interface{}, w *http.ResponseWriter) {
	user := i.(*AddFriendData)
	reg := IAddFriend{
		Retcode:  200,
		Msg:      "ok",
		Datetime: 10,
		Data:     AddFriendData{Id: user.Id, Name: user.Name},
	}
	encode := json.NewEncoder(*w)
	encode.Encode(reg)
}
