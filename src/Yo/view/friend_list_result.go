package view

import (
	"encoding/json"
	"net/http"
	"yo/module"
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
	var lstSendData GetFriendListData
	lstUser := *(i.(*[]module.UserData))
	lstSendData.Count = len(lstUser)

	for _, value := range lstUser {
		var u User
		u.Name = value.Name
		u.Id = value.Id
		lstSendData.Friend = append(lstSendData.Friend, u)
	}

	reg := IGetFriendList{
		Retcode:  200,
		Msg:      "ok",
		Datetime: 10,
		Data:     lstSendData,
	}

	encode := json.NewEncoder(*w)
	encode.Encode(reg)
}
