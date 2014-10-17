package yo

import (
	"httprouter"
	"net/http"
	"restcontrol"
	"yo/module"
)

type HttpServer struct {
	Server *module.ModuleServer
	Router *httprouter.HttpRouter
}

/*
函数名：初始化

返回值：error 错误码

success

complaxinit 重复初始化
*/
func (h *HttpServer) Init() string {
	if h.Server != nil {
		return "complaxinit"
	}

	if h.Router != nil {
		return "complaxinit"
	}

	h.Server = new(module.ModuleServer)
	h.Router = new(httprouter.HttpRouter)
	return h.Router.Init()
}

/*
函数名：增加一个 control

返回值：error 错误码

success

uninit 未初始化

isexist 重复绑定
*/
func (h *HttpServer) AddControl(szPart string, control restcontrol.RESTControl) string {
	if h.Server == nil || h.Router == nil {
		return "uninit"
	}

	return h.Router.AddControl(szPart, control)
}

/*
函数名：启动服务

返回值：error 错误码

success

uninit 未初始化

httperr golang http 服务内部错误
*/
func (h *HttpServer) Start(szPort string) string {
	if h.Router == nil || h.Server == nil {
		return "uninit"
	}

	err := http.ListenAndServe(szPort, h.Router)
	if err == nil {
		return "success"
	} else {
		return err.Error()
	}
}
