package control

import (
	"base"
	"net/http"
	"yo/module"
	"yo/view"
)

type UserControl struct {
	BaseControl
}

/*
输出的错误值有:

missusername

serverfaile
*/
// 新建一个用户
func (uc *UserControl) Post(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Post")

	var szUserName string
	var ok bool
	var err string
	var user *module.UserData
	var data view.RegUserData

	szUserName, ok = uc.tbParam["user"]
	if ok == false {
		err = "missusername"
		goto SEND
	}

	err, user = uc.svr.RegistUser(szUserName)
	data = view.RegUserData{Userid: user.Id, Username: user.Name}
SEND:
	if err == "success" {
		err = uc.vm.DoRender("reguser", uc.szViewType, &data, w)
		if err != "success" {
			goto SEND
		}
	} else {
		err = uc.vm.DoRender("error", uc.szViewType, err, w)
		if err != "success" {
			Base.PrintErr("can not find view type :error.json")
		}
	}
}
