package control

import (
	"base"
	"httprouter"
	"net/http"
	"yo"
	"yo/module"
	"yo/view"
)

type BaseControl struct {
	tbParam    map[string]string
	vm         *view.ViewManager
	svr        *module.ModuleServer
	szViewType string
}

func (bc *BaseControl) Init(w *http.ResponseWriter, r *http.Request, tbParam map[string]string) (err string) {
	Base.PrintLog("Init")
	bc.tbParam = tbParam

	var ok bool
	bc.szViewType, ok = bc.tbParam[httprouter.ViewTypeName]
	if ok == false {
		err = "missviewtype"
		goto SEND
	}

	err, bc.vm = yo.GetViewManager()
	if err != "success" {
		err = "viewmanagerfail"
		goto SEND
	}

	err, bc.svr = yo.GetModuleServer()
	if err != "success" {
		err = "moduleserverfail"
		goto SEND
	}

SEND:
	if err != "success" {
		if err != "missviewtype" && err != "viewmanagerfail" {
			err = bc.vm.DoRender("error", bc.szViewType, err, w)
		}

		Base.PrintErr("can not find view type :error.json, the error is " + err)
	}

	return err
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
