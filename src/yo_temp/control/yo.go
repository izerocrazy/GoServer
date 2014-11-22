package control

import (
	"base"
	"net/http"
	"strconv"
	"yo/module"
	"yo/view"
)

type YoControl struct {
	BaseControl
}

// 获取未读的 yo
func (yc *YoControl) Get(w *http.ResponseWriter, r *http.Request) {
	var err string
	var szUserId string
	var userId int
	var ok bool
	var lstYO []module.MsgInfo
	var errEr error
	var lstData view.GetYOData

	szUserId, ok = yc.tbParam["usersessionid"]
	if ok == false {
		err = "emptyuserid"
		goto SEND
	}

	userId, errEr = strconv.Atoi(szUserId)
	if errEr != nil {
		err = "erroruserid"
		goto SEND
	}

	err, lstYO = yc.svr.GetYO(userId)
	if err != "success" {
		goto SEND
	}

	lstData.Count = len(lstYO)
	for _, value := range lstYO {
		var y view.YOMsg
		err, y.From = yc.svr.GetUserName(value.SenderId)
		if err == "success" {
			lstData.Msgs = append(lstData.Msgs, y)
		} /*else {
			// reg = makeError(err)
			// break
		}*/
	}
	err = "success"

SEND:
	if err == "success" {
		err = yc.vm.DoRender("getyo", yc.szViewType, &lstData, w)
		if err != "success" {
			goto SEND
		}
	} else {
		err = yc.vm.DoRender("error", yc.szViewType, err, w)
		if err != "success" {
			Base.PrintErr("can not find view type :error.json " + err)
		}
	}
}

// 发送 yo
func (yc *YoControl) Post(w *http.ResponseWriter, r *http.Request) {
	var err string
	var user *module.UserData

	var ok bool
	var szFriendName string
	var szUserId string
	var userId int
	var geterId int
	var errEr error

	szFriendName, ok = yc.tbParam["yo"]
	if ok == false {
		err = "emptyfriendname"
		goto SEND
	}

	szUserId, ok = yc.tbParam["usersessionid"]
	if ok == false {
		err = "emptyuserid"
		goto SEND
	}

	userId, errEr = strconv.Atoi(szUserId)
	if errEr != nil {
		err = "erroruserid"
		goto SEND
	}

	user = yc.svr.GetUserByName(szFriendName)
	geterId = user.Id
	err = yc.svr.SendYO(userId, geterId)

SEND:
	if err == "success" {
		err = yc.vm.DoRender("sendyo", yc.szViewType, nil, w)
		if err != "success" {
			goto SEND
		}
	} else {
		err = yc.vm.DoRender("error", yc.szViewType, err, w)
		if err != "success" {
			Base.PrintErr("can not find view type :error.json " + err)
		}
	}
}
