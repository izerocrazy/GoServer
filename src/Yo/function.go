package yo

import (
	"base"
	"restcontrol"
	"yo/module"
	"yo/view"
)

var s *HttpServer
var ms *module.ModuleServer
var vm *view.ViewManager

/*
函数名：初始化

返回值：error 错误码

success

httpservererror

viewmanagererror

complaxinit 重复初始化
*/
func Init() string {
	if s == nil {
		s = new(HttpServer)
		err := s.Init()
		if err != "success" {
			return "httpservererror"
		}
	}

	if ms == nil {
		ms = new(module.ModuleServer)
	}

	if vm == nil {
		vm = new(view.ViewManager)
		err := vm.Init()
		if err != "success" {
			return "viewmanagererror"
		}
	}

	return "success"
}

/*
函数名：增加一个 control

返回值：error 错误码

success

uninit 未初始化

isexist 重复绑定
*/
func AddControl(szPath string, control restcontrol.RESTControl) string {
	if s == nil {
		return "uninit"
	}

	return s.AddControl(szPath, control)
}

/*
函数名：启动服务

返回值：error 错误码

success

uninit 未初始化

httperr golang http 服务内部错误
*/
func StartServer() string {
	if s == nil {
		return "uninit"
	}

	return s.Start(":8080")
}

/*
函数名：得到服务

返回值：error 错误码

success

uninit 未初始化
*/
func GetModuleServer() (err string, svr *module.ModuleServer) {
	if s == nil || ms == nil {
		return "uninit", nil
	}

	return "success", ms
}

/*
函数名：得到view

返回值：error 错误码

success

uninit 未初始化
*/
func GetViewManager() (err string, vm2 *view.ViewManager) {
	if s == nil || vm == nil {
		return "uninit", nil
	}

	vm2 = vm
	return "success", vm2
}
