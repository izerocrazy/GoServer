package view

import (
	"net/http"
	"reflect"
	"reflectmap"
)

type ViewManager struct {
	RenderMap *reflectmap.ReflectMap
}

// 错误码
// success
// complaxinit 重复初始化
// 其余参见 reflectmap.Init 以及 regsterView
func (vm *ViewManager) Init() (err string) {
	if vm.RenderMap == nil {
		vm.RenderMap = new(reflectmap.ReflectMap)
		err = vm.RenderMap.Init()
		if err != "success" {
			goto ERROR
		}

		err = vm.regsterView()
		if err != "success" {
			goto ERROR
		}
	} else {
		err = "complaxinit"
		goto ERROR
	}

	err = "success"
ERROR:
	return err
}

// 错误码
// success
// uninit 未初始化
// 其余参见 reflectmap.Init 以及 AddRenderMapWithType
func (vm *ViewManager) regsterView() string {
	if vm.RenderMap == nil {
		return "uninit"
	}

	var rr RegUserResult
	var re ErrResult
	err := vm.AddRenderMapWithType("reguser", "json", &rr)
	if err != "success" {
		goto ERROR
	}

	err = vm.AddRenderMapWithType("error", "json", &re)
	if err != "success" {
		goto ERROR
	}

	err = "success"
ERROR:
	return err
}

// 错误码
// success
// uninit
// 其余参见 reflectmap.Add()
func (vm *ViewManager) AddRenderMap(szName string, r Render) string {
	if vm.RenderMap == nil {
		return "uninit"
	}

	return vm.RenderMap.Add(szName, r)
}

func checkType(szType string) bool {
	if szType != "json" && szType != "xml" && szType != "html" {
		return false
	}

	return true
}

// 错误码
// success
// uninit
// 其余参见 AddRenderMap
func (vm *ViewManager) AddRenderMapWithType(szName string, szType string, r Render) string {
	if szName == "" {
		return "emptystring"
	}

	if checkType(szType) == false {
		return "errortype"
	}

	if vm.RenderMap == nil {
		return "uninit"
	}

	return vm.RenderMap.Add(szName, r)
}

// 错误码
// success
// uninit
// 其余参见 reflectmap.Add()
func (vm *ViewManager) GetRender(szName string) (err string, v reflect.Value) {
	if vm.RenderMap == nil {
		return "uninit", v
	}
	return vm.RenderMap.New(szName)
}

// 错误码
// success
// emptystring 名字为空
// errortype 类型错误
// 其余参加 GetRender
// 此接口预留给以后做 fatory 相关的内容
func (vm *ViewManager) GetRenderByType(szName string, szType string) (err string, v reflect.Value) {
	if vm.RenderMap == nil {
		return "uninit", v
	}

	if szName == "" {
		return "emptystring", v
	}

	if checkType(szType) == false {
		return "errortype", v
	}

	szAllName := szName + "." + szType

	return vm.GetRender(szAllName)
}

// 错误码参考 GetRenderByType
func (vm *ViewManager) DoRender(szName string, szType string, i interface{}, w *http.ResponseWriter) (err string) {
	err, r := vm.GetRenderByType(szName, szType)
	if err == "success" {
		render := r.MethodByName("Render")
		tbParam := make([]reflect.Value, 2)
		tbParam[0] = reflect.ValueOf(i)
		tbParam[1] = reflect.ValueOf(w)
		render.Call(tbParam)
	}

	return err
}
