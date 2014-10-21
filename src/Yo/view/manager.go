package view

import (
	"reflect"
	"reflectmap"
)

type ViewManager struct {
	RenderMap *reflectmap.ReflectMap
}

// 错误码
// success
// complaxinit 重复初始化
func (vm *ViewManager) Init() string {
	if vm.RenderMap == nil {
		vm.RenderMap = new(reflectmap.ReflectMap)
		return vm.RenderMap.Init()
	}

	return "complaxinit"
}

// 错误码
// success
// uninit
// isexist 这个字符串已经有了对应的 reflect.type
func (vm *ViewManager) AddRenderMap(szName string, r Render) string {
	if vm.RenderMap == nil {
		return "uninit"
	}

	return vm.RenderMap.Add(szName, r)
}

// 错误码
// success
// regempty 这个字符串没有了对应的 reflect.type
// uninit
func (vm *ViewManager) GetRender(szName string) (err string, v reflect.Value) {
	if vm.RenderMap == nil {
		return "uninit", v
	}
	return vm.RenderMap.New(szName)
}
