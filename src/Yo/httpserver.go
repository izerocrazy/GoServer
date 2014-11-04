package yo

import (
	"httprouter"
	"net/http"
	"restcontrol"
	"yo/view"
)

type HttpServer struct {
	Router *httprouter.HttpRouter
}

/*
函数名：初始化

返回值：error 错误码

success

complaxinit 重复初始化

其余参考 HttpRouter.Init() 以及 ViewManager.Init()
*/
func (h *HttpServer) Init() string {
	if h.Router != nil {
		return "complaxinit"
	}

	h.Router = new(httprouter.HttpRouter)
	err := h.Router.Init()
	if err != "success" {
		return err
	}

	return "success"
}

/*
函数名：增加一个 control

返回值：error 错误码

success

uninit 未初始化

其余参考 HttpRouter.AddControl()
*/
func (h *HttpServer) AddControl(szPart string, control restcontrol.RESTControl) string {
	if h.Router == nil {
		return "uninit"
	}

	return h.Router.AddControl(szPart, control)
}

/*
函数名：增加一个 view render

返回值：error 错误码

success

uninit 未初始化

其余参考 ViewMananger.AddRenderMapWithType
*/
func (h *HttpServer) addViewWithType(szName string, szType string, view view.Render) string {
	if h.Router == nil {
		return "uninit"
	}

	err, vm := GetViewManager()
	if vm == nil || err != "success" {
		return err
	}

	return vm.AddRenderMapWithType(szName, szType, view)
}

/*
函数名：启动服务

返回值：error 错误码

success

uninit 未初始化

httperr golang http 服务内部错误
*/
func (h *HttpServer) Start(szPort string) string {
	if h.Router == nil {
		return "uninit"
	}

	err := http.ListenAndServe(szPort, h.Router)
	if err == nil {
		return "success"
	} else {
		return err.Error()
	}
}
