package control

import (
	"base"
	"net/http"
)

type UserControl struct {
}

func (uc *UserControl) Init(w *http.ResponseWriter, r *http.Request, tbParam map[string]string) {
	Base.PrintLog("Init")
}

// 得到一个用户的信息
func (uc *UserControl) Get(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Get")
}

// 新建一个用户
func (uc *UserControl) Post(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Post")
}

// 修改一个用户信息
func (uc *UserControl) Put(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Put")
}

// 删除
func (uc *UserControl) Delete(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Delete")
}
