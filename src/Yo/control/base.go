package control

import (
	"base"
	"net/http"
)

type BaseControl struct {
	tbParam map[string]string
}

func (bc *BaseControl) Init(w *http.ResponseWriter, r *http.Request, tbParam map[string]string) {
	Base.PrintLog("Init")
	bc.tbParam = tbParam
}

func (bc *BaseControl) Get(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Get")
}

func (bc *BaseControl) Post(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Post")
}

func (bc *BaseControl) Put(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Put")
}

func (bc *BaseControl) Delete(w *http.ResponseWriter, r *http.Request) {
	Base.PrintLog("Delete")
}
