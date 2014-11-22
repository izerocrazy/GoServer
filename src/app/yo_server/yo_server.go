package main

import (
	"base"
	"yo"
	"yo/control"
)

func main() {
	// HtmlServer := http.FileServer(http.Dir("."))
	// http.Handle("/", HtmlServer)
	// http.HandleFunc("/test", testFun)

	// http.HandleFunc("/user/reg", RegUser)
	// http.HandleFunc("/friend/add", AddFriend)
	// http.HandleFunc("/friend/list", GetFriendList)
	// http.HandleFunc("/yo/sendyo", SendYO)
	// http.HandleFunc("/yo/getyo", GetYO)

	// // 缺少一个默认的 404

	// err := http.ListenAndServe(":8000", nil)
	// Base.CheckErr(err)

	yo.Init()
	yo.AddControl("/user", &control.UserControl{})
	yo.AddControl("/", &control.UserControl{})
	yo.AddControl("/friend", &control.FriendControl{})
	yo.AddControl("/yo", &control.YoControl{})
	szCode := yo.StartServer()
	if szCode != "success" {
		Base.PrintErrExit(szCode)
	}
}
