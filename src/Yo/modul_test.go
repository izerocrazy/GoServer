package yo

import (
	"testing"
)

func TestRegistUser(t *testing.T) {
	t.Log("test RegistUser")
	var s Server
	// 第一次肯定成功
	err, u := s.RegistUser("name1")
	if err != "success" {
		t.Log("Regist User Error: for Regist User name1")
		t.Fail()
	}

	// 第二次不重名，也必须成功
	err, u = s.RegistUser("name2")
	if err != "success" {
		t.Log("Regist User Error: for Regist User name2")
		t.Fail()
	}

	// 第二次重名，必须失败
	err, u = s.RegistUser("name1")
	if err != "nameexist" && u != nil {
		t.Log("Regist User Error: for Regist User name1 again")
		t.Fail()
	}
}

func TestAddFriend(t *testing.T) {
	t.Log("TestAddFriend")
	var s Server
	err1, u1 := s.RegistUser("user1")
	if err1 != "success" {
		t.Log("Regist User Error")
		t.Fail()
	}
	err2, _ := s.RegistUser("user2")
	if err2 != "success" {
		t.Log("Regist User Error")
		t.Fail()
	}

	// 第一次必须成功
	err := s.AddFriend(u1.Id, "user2")
	if err != "success" {
		t.Log("Add Friend Error: User1 add User2 fail")
		t.Fail()
	}

	// 自己添加自己，必须失败
	err = s.AddFriend(u1.Id, "user1")
	if err != "idnameissameone" {
		t.Log("Add Friend Error: self add should be fail")
		t.Fail()
	}

	// 用一个空的 Id，去加用户，必须失败
	err = s.AddFriend(-1, "user3")
	if err != "iduserempty" {
		t.Log("Add Friend Error: User id is error")
		t.Fail()
	}

	// 用一个真实 Id，去加一个空用户，必须失败
	err = s.AddFriend(u1.Id, "user3")
	if err != "nameuserempty" {
		t.Log("Add Friend Error: User name is error")
		t.Fail()
	}

	// 增加一个已是好友的用户，必须失败
	err = s.AddFriend(u1.Id, "user2")
	if err != "alreadyfriend" {
		t.Log("Add Friend Error: already friend")
		t.Fail()
	}
}

func TestGetFriendList(t *testing.T) {
	t.Log("test GetFriendList")
	var s Server
	_, user1 := s.RegistUser("user1")

	// 取一个不存在用户的 contact list
	err, lstContact := s.GetFriendList(-1)
	if err != "emptyuser" {
		t.Log("Get Friend Error: user should be empty")
		t.Fail()
	}

	// 测试列表为空
	err, lstContact = s.GetFriendList(user1.Id)
	if len(lstContact) != 0 {
		t.Log("Get Friend Error: lst should be empty")
		t.Fail()
	}

	// 测试列表有一个
	s.RegistUser("user2")
	err = s.AddFriend(user1.Id, "user2")
	if err != "success" {
		t.Log("Get Friend List Error: Add Friend error: " + err)
		t.Fail()
	}
	err, lstContact = s.GetFriendList(user1.Id)
	if len(lstContact) != 1 {
		t.Log("Get Friend Error: should have one friend")
		t.Fail()
	}

	// 测试列表有两个
	s.RegistUser("user3")
	err = s.AddFriend(user1.Id, "user3")
	if err != "success" {
		t.Log("Get Friend List Error: Add Friend error: " + err)
		t.Fail()
	}
	err, lstContact = s.GetFriendList(user1.Id)
	if len(lstContact) != 2 {
		t.Log("Get Friend Error: should have two friend")
		t.Fail()
	}
}

func TestSendYO(t *testing.T) {
	t.Log("test SendYO")
}

func TestGetYO(t *testing.T) {
	t.Log("test GetYO")
}
