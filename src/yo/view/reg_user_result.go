package view

import (
	"encoding/json"
	"net/http"
	"yo/module"
)

type RegUserData struct {
	Userid   int    `json:"userid"`
	Username string `json:"username"`
}

type IRegUser struct {
	Retcode  int         `json:"retcode"`
	Msg      string      `json:"msg"`
	Datetime int         `json:"datetime"`
	Data     RegUserData `json:"data"`
}

type RegUserResult struct {
}

func (rr *RegUserResult) render(user *module.UserData) (reg interface{}) {
	// 设定 cookie
	// cookiename := "username_" + szUserName
	// cookieid := "userid+" + user.Id
	// cookie := http.Cookie{Name: "userid", Value: fmt.Sprintf("%d", user.Id), Path: "/"}
	// http.SetCookie(w, &cookie)

	return reg
}

func (rr *RegUserResult) Render(i interface{}, w *http.ResponseWriter) {
	// 这里不处理强制转换失败了，因为如果出现错误，也需要重新编译
	data := *(i.(*RegUserData))
	reg := IRegUser{
		Retcode:  200,
		Msg:      "ok",
		Datetime: 10,
		Data:     data,
	}
	encode := json.NewEncoder(*w)
	encode.Encode(reg)
}
