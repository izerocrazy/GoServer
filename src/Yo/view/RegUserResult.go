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

// 思考怎么实现数据上的继承
type IRegUser struct {
	Retcode  int         `json:"retcode"`
	Msg      string      `json:"msg"`
	Datetime int         `json:"datetime"`
	Data     RegUserData `json:"data"`
}

type RegUserResult struct {
}

func (rr *RegUserResult) rander(err string, user *module.UserData) (reg interface{}) {
	if err == "success" {
		// 设定 cookie
		// cookiename := "username_" + szUserName
		// cookieid := "userid+" + user.Id
		// cookie := http.Cookie{Name: "userid", Value: fmt.Sprintf("%d", user.Id), Path: "/"}
		// http.SetCookie(w, &cookie)

		reg = IRegUser{
			Retcode:  200,
			Msg:      "ok",
			Datetime: 10,
			Data:     RegUserData{Userid: user.Id, Username: user.Name},
		}
	} else {
		reg = MakeError(err)
	}

	return reg
}

func (rr *RegUserResult) Rander(err string, user *module.UserData, w *http.ResponseWriter) {
	reg := rr.rander(err, user)
	encode := json.NewEncoder(*w)
	encode.Encode(reg)
}
