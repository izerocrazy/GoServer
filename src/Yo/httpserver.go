package yo

import (
	"httprouter"
	"net/http"
	"restcontrol"
	"yo/module"
	"yo/view"
)

type HttpServer struct {
	MServer     *module.ModuleServer
	Router      *httprouter.HttpRouter
	ViewManager *view.ViewManager
}

/*
函数名：初始化

返回值：error 错误码

success

complaxinit 重复初始化

其余参考 HttpRouter.Init() 以及 ViewManager.Init()
*/
func (h *HttpServer) Init() string {
	if h.ViewManager != nil || h.MServer != nil || h.Router != nil {
		return "complaxinit"
	}

	h.MServer = new(module.ModuleServer)
	h.Router = new(httprouter.HttpRouter)
	h.ViewManager = new(view.ViewManager)
	err := h.Router.Init()
	if err != "success" {
		return err
	}

	err = h.ViewManager.Init()
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
	if h.MServer == nil || h.Router == nil || h.ViewManager == nil {
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
	if h.MServer == nil || h.Router == nil || h.ViewManager == nil {
		return "uninit"
	}

	return h.ViewManager.AddRenderMapWithType(szName, szType, view)
}

/*
函数名：启动服务

返回值：error 错误码

success

uninit 未初始化

httperr golang http 服务内部错误
*/
func (h *HttpServer) Start(szPort string) string {
	if h.Router == nil || h.MServer == nil {
		return "uninit"
	}

	err := http.ListenAndServe(szPort, h.Router)
	if err == nil {
		return "success"
	} else {
		return err.Error()
	}
}
