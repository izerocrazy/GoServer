package yo

import (
	"fmt"
	"testing"
)

func TestRegistUser(t *testing.T) {
	t.Log("test RegistUser")
	var s Server
	// 第一次肯定成功
	err, u := s.RegistUser("name1")
	if err != "success" {
		t.Log("Regist User Error: for Regist User name1", err)
		t.FailNow()
	}

	// 第二次不重名，也必须成功
	err, u = s.RegistUser("name2")
	if err != "success" {
		t.Log("Regist User Error: for Regist User name2")
		t.FailNow()
	}

	// 第二次重名，必须失败
	err, u = s.RegistUser("name1")
	if err != "nameexist" && u != nil {
		t.Log("Regist User Error: for Regist User name1 again")
		t.FailNow()
	}
}

func TestAddFriend(t *testing.T) {
	t.Log("TestAddFriend")
	var s Server
	err1, u1 := s.RegistUser("user1")
	if err1 != "success" {
		t.Log("Regist User Error")
		t.FailNow()
	}
	err2, _ := s.RegistUser("user2")
	if err2 != "success" {
		t.Log("Regist User Error")
		t.FailNow()
	}

	// 第一次必须成功
	err := s.AddFriend(u1.Id, "user2")
	if err != "success" {
		t.Log("Add Friend Error: User1 add User2 fail", err)
		t.FailNow()
	}

	// 自己添加自己，必须失败
	err = s.AddFriend(u1.Id, "user1")
	if err != "idnameissameone" {
		t.Log("Add Friend Error: self add should be fail")
		t.FailNow()
	}

	// 用一个空的 Id，去加用户，必须失败
	err = s.AddFriend(3, "user3")
	if err != "iduserempty" {
		t.Log("Add Friend Error: User id is error")
		t.FailNow()
	}

	// 用一个真实 Id，去加一个空用户，必须失败
	err = s.AddFriend(u1.Id, "user3")
	if err != "nameuserempty" {
		t.Log("Add Friend Error: User name is error")
		t.FailNow()
	}

	// 增加一个已是好友的用户，必须失败
	err = s.AddFriend(u1.Id, "user2")
	if err != "alreadyfriend" {
		t.Log("Add Friend Error: already friend")
		t.FailNow()
	}
}

func TestGetFriendList(t *testing.T) {
	t.Log("test GetFriendList")
	var s Server
	_, user1 := s.RegistUser("user1")

	// 取一个不存在用户的 contact list
	err, lstContact := s.GetFriendList(2)
	if err != "emptyuser" {
		t.Log("Get Friend Error: user should be empty")
		t.FailNow()
	}

	// 测试列表为空
	err, lstContact = s.GetFriendList(user1.Id)
	if len(lstContact) != 0 {
		t.Log("Get Friend Error: lst should be empty")
		t.FailNow()
	}

	// 测试列表有一个
	s.RegistUser("user2")
	err = s.AddFriend(user1.Id, "user2")
	if err != "success" {
		t.Log("Get Friend List Error: Add Friend error: " + err)
		t.FailNow()
	}
	err, lstContact = s.GetFriendList(user1.Id)
	if len(lstContact) != 1 {
		t.Log("Get Friend Error: should have one friend")
		t.FailNow()
	}

	if lstContact[0].Id != 1 || lstContact[0].Name != "user2" {
		t.Log("Get Friend Error: friend data is error")
		fmt.Println(lstContact)
		t.FailNow()
	}

	// 测试列表有两个
	s.RegistUser("user3")
	err = s.AddFriend(user1.Id, "user3")
	if err != "success" {
		t.Log("Get Friend List Error: Add Friend error: " + err)
		t.FailNow()
	}
	err, lstContact = s.GetFriendList(user1.Id)
	if len(lstContact) != 2 {
		t.Log("Get Friend Error: should have two friend")
		t.FailNow()
	}
}

func TestSendYO(t *testing.T) {
	t.Log("test SendYO")
	var s Server
	// 新增两个用户
	err1, user1 := s.RegistUser("user1")
	if err1 != "success" {
		t.Log("Send YO Error: Create user1 error" + err1)
		t.FailNow()
	}
	err2, user2 := s.RegistUser("user2")
	if err2 != "success" {
		t.Log("Send YO Error: Create user2 error" + err1)
		t.FailNow()
	}

	// 参数检查
	err := s.SendYO(2, user2.Id)
	if err != "emptysender" {
		t.Log("Send YO Error: Empty Sender" + err)
		t.FailNow()
	}

	err = s.SendYO(user1.Id, 2)
	if err != "emptygeter" {
		t.Log("Send YO Error: Empty Geter" + err)
		t.FailNow()
	}

	err = s.SendYO(user1.Id, user1.Id)
	if err != "sendergetersame" {
		t.Log("Send YO Error: Same User" + err)
		t.FailNow()
	}

	// 不加好友不能够发送 yo
	err = s.SendYO(user1.Id, user2.Id)
	if err != "nofriend" {
		t.Log("Send YO Error: User1 and User2 is no friend" + err)
		t.FailNow()
	}

	err = s.AddFriend(user1.Id, "user2")
	if err != "success" {
		t.Log("Send YO Error: User1 add User2 fail" + err)
		t.FailNow()
	}

	// 加好友后发送 yo 成功
	err = s.SendYO(user1.Id, user2.Id)
	if err != "success" {
		t.Log("Send YO Error: user1 send YO to user2" + err)
		t.FailNow()
	}
}

func TestGetYO(t *testing.T) {
	t.Log("test GetYO")
	var s Server
	// 新增两个用户
	err1, user1 := s.RegistUser("user1")
	if err1 != "success" {
		t.Log("Get YO Error: Create user1 error" + err1)
		t.FailNow()
	}
	err2, user2 := s.RegistUser("user2")
	if err2 != "success" {
		t.Log("Get YO Error: Create user2 error" + err1)
		t.FailNow()
	}

	// 参数检查
	err, _ := s.GetYO(2)
	if err != "emptyuser" {
		t.Log("Get YO Error: Get Empty User" + err)
		t.FailNow()
	}

	// 加为好友，发送 yo 成功
	err = s.AddFriend(user1.Id, "user2")
	if err != "success" {
		t.Log("Get YO Error: User1 add User2 fail" + err)
		t.FailNow()
	}

	err = s.SendYO(user1.Id, user2.Id)
	if err != "success" {
		t.Log("Get YO Error: user1 send YO to user2" + err)
		t.FailNow()
	}

	err = s.SendYO(user1.Id, user2.Id)
	if err != "success" {
		t.Log("Get YO Error: user1 send YO to user2" + err)
		t.FailNow()
	}

	errlst, lst := s.GetYO(user2.Id)
	if errlst != "success" && len(lst) != 2 {
		t.Log("Get YO Error: User2 YO list err" + err)
		t.FailNow()
	}

	if lst[0].SenderId != user1.Id || lst[0].GeterId != user2.Id {
		t.Log("Get YO Error: YO Data Err")
		fmt.Println(lst[0])
		t.FailNow()
	}
}
