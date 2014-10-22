package control

import (
	"base"
	"httprouter"
	"net/http"
	"yo"
	"yo/module"
	"yo/view"
)

type UserControl struct {
	tbParam map[string]string
}

func (uc *UserControl) Init(w *http.ResponseWriter, r *http.Request, tbParam map[string]string) {
	Base.PrintLog("Init")

	uc.tbParam = tbParam
}

// 得到一个用户的信息
func (uc *UserControl) Get(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Get")
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
	var szViewType string
	var err string
	var user *module.UserData
	var ok bool
	var vm *view.ViewManager

	err, svr := yo.GetModuleServer()
	if err != "success" {
		err = "moduleserverfail"
		goto SEND
	}

	err, vm = yo.GetViewManager()
	if err != "success" {
		err = "viewmanagerfail"
		goto SEND
	}

	szUserName, ok = uc.tbParam["user"]
	if ok == false {
		err = "missusername"
		goto SEND
	}

	szViewType, ok = uc.tbParam[httprouter.ViewTypeName]
	if ok == false {
		err = "missviewtype"
		goto SEND
	}

	err, user = svr.RegistUser(szUserName)
SEND:
	if err == "success" {
		err = vm.DoRender("reguser", szViewType, user, w)
		if err != "success" {
			goto SEND
		}
	} else {
		err = vm.DoRender("error", szViewType, err, w)
		if err != "success" {
			Base.PrintErr("can not find view type :error.json")
		}
	}
}

// 修改一个用户信息
func (uc *UserControl) Put(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Put")
}

// 删除
func (uc *UserControl) Delete(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Delete")
}
