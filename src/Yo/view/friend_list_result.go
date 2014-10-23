package view

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type GetFriendListData struct {
	Count  int    `json:"count"`
	Friend []User `json:"friend"`
}

// 思考怎么实现数据上的继承
type IGetFriendList struct {
	Retcode  int               `json:"retcode"`
	Msg      string            `json:"msg"`
	Datetime int               `json:"datetime"`
	Data     GetFriendListData `json:"data"`
}

type FriendListResult struct {
}

func (fr *FriendListResult) Render(i interface{}, w *http.ResponseWriter) {
	lstSendData := *(i.(*GetFriendListData))

	reg := IGetFriendList{
		Retcode:  200,
		Msg:      "ok",
		Datetime: 10,
		Data:     lstSendData,
	}

	encode := json.NewEncoder(*w)
	encode.Encode(reg)
}
