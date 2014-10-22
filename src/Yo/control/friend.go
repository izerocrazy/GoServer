package control

import (
	"base"
	"httprouter"
	"net/http"
	"strconv"
	"yo"
	"yo/module"
	"yo/view"
)

type FriendControl struct {
	BaseControl
}

func (fc *FriendControl) Get(w *http.ResponseWriter, r *http.Request) {
	var err string
	var vm *view.ViewManager
	var svr *module.ModuleServer
	var szUserId string
	var szViewType string
	var userId int
	var ok bool
	var lstUser []module.UserData
	var errEr error

	szViewType, ok = fc.tbParam[httprouter.ViewTypeName]
	if ok == false {
		szViewType = "json"
		err = "missviewtype"
		goto SEND
	}

	err, svr = yo.GetModuleServer()
	if err != "success" {
		err = "moduleserverfail"
		goto SEND
	}

	err, vm = yo.GetViewManager()
	if err != "success" {
		err = "viewmanagerfail"
		goto SEND
	}

	szUserId, ok = fc.tbParam["usersessionid"]
	if ok == false {
		err = "emptyuserid"
		goto SEND
	}

	userId, errEr = strconv.Atoi(szUserId)
	if errEr != nil {
		err = "erroruserid"
		goto SEND
	}

	err, lstUser = svr.GetFriendList(userId)
	if err != "success" {
		goto SEND
	}

SEND:
	if err == "success" {
		err = vm.DoRender("friendlist", szViewType, &lstUser, w)
		if err != "success" {
			goto SEND
		}
	} else {
		err = vm.DoRender("error", szViewType, err, w)
		if err != "success" {
			Base.PrintErr("can not find view type :error.json " + err)
		}
	}
}

// Post ：添加用户
func (fc *FriendControl) Post(w *http.ResponseWriter, r *http.Request) {
	var err string
	var user *module.UserData
	var svr *module.ModuleServer
	var vm *view.ViewManager

	var ok bool
	var szFriendName string
	var szViewType string
	var szUserId string
	var userId int
	var errEr error

	var va view.AddFriendData

	szViewType, ok = fc.tbParam[httprouter.ViewTypeName]
	if ok == false {
		err = "missviewtype"
		goto SEND
	}

	err, svr = yo.GetModuleServer()
	if err != "success" {
		err = "moduleserverfail"
		goto SEND
	}

	err, vm = yo.GetViewManager()
	if err != "success" {
		err = "viewmanagerfail"
		goto SEND
	}

	szFriendName, ok = fc.tbParam["friend"]
	if ok == false {
		err = "emptyfriendname"
		goto SEND
	}

	szUserId, ok = fc.tbParam["usersessionid"]
	if ok == false {
		err = "emptyuserid"
		goto SEND
	}

	userId, errEr = strconv.Atoi(szUserId)
	if errEr != nil {
		err = "erroruserid"
		goto SEND
	}

	err = svr.AddFriend(userId, szFriendName)
	if err != "success" {
		goto SEND
	}

	user = svr.GetUserByName(szFriendName)
	va.Id = user.Id
	va.Name = user.Name

SEND:
	if err == "success" {
		err = vm.DoRender("addfriend", szViewType, &va, w)
		if err != "success" {
			goto SEND
		}
	} else {
		err = vm.DoRender("error", szViewType, err, w)
		if err != "success" {
			Base.PrintErr("can not find view type :error.json " + err)
		}
	}
}
