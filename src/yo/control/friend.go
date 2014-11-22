package control

import (
	"base"
	"net/http"
	"strconv"
	"yo/module"
	"yo/view"
)

type FriendControl struct {
	BaseControl
}

func (fc *FriendControl) Get(w *http.ResponseWriter, r *http.Request) {
	var err string
	var szUserId string
	var userId int
	var ok bool
	var lstUser []module.UserData
	var lstSendData view.GetFriendListData
	var errEr error

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

	err, lstUser = fc.svr.GetFriendList(userId)
	if err != "success" {
		goto SEND
	}

	lstSendData.Count = len(lstUser)
	for _, value := range lstUser {
		var u view.User
		u.Name = value.Name
		u.Id = value.Id
		lstSendData.Friend = append(lstSendData.Friend, u)
	}

SEND:
	if err == "success" {
		err = fc.vm.DoRender("friendlist", fc.szViewType, &lstSendData, w)
		if err != "success" {
			goto SEND
		}
	} else {
		err = fc.vm.DoRender("error", fc.szViewType, err, w)
		if err != "success" {
			Base.PrintErr("can not find view type :error.json " + err)
		}
	}
}

// Post ：添加用户
func (fc *FriendControl) Post(w *http.ResponseWriter, r *http.Request) {
	var err string
	var user *module.UserData

	var ok bool
	var szFriendName string
	var szUserId string
	var userId int
	var errEr error

	var va view.AddFriendData

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

	err = fc.svr.AddFriend(userId, szFriendName)
	if err != "success" {
		goto SEND
	}

	user = fc.svr.GetUserByName(szFriendName)
	va.Id = user.Id
	va.Name = user.Name

SEND:
	if err == "success" {
		err = fc.vm.DoRender("addfriend", fc.szViewType, &va, w)
		if err != "success" {
			goto SEND
		}
	} else {
		err = fc.vm.DoRender("error", fc.szViewType, err, w)
		if err != "success" {
			Base.PrintErr("can not find view type :error.json " + err)
		}
	}
}
