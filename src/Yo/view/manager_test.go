package view

import (
	"net/http"
	"testing"
	"yo/module"
)

func TestInit(t *testing.T) {
	var vm ViewManager

	err := vm.Init()
	if err != "success" {
		t.Log("Test View Mananger Init Err", err)
		t.FailNow()
	}

	err = vm.Init()
	if err != "complaxinit" {
		t.Log("Test View Mananger Init Err", err)
		t.FailNow()
	}
}

type TestRender struct{}

func (tr *TestRender) Render(err string, user *module.UserData, w *http.ResponseWriter) {

}

func TestAddRenderMap(t *testing.T) {
	var vm ViewManager

	var i TestRender

	err := vm.AddRenderMap("x_json", &i)
	if err != "uninit" {
		t.Log("Test View Mananger Add Render Map Err", err)
		t.FailNow()
	}

	err = vm.Init()
	if err != "success" {
		t.Log("Test View Mananger Add Render Map Err", err)
		t.FailNow()
	}

	err = vm.AddRenderMap("x_json", &i)
	if err != "success" {
		t.Log("Test View Mananger Add Render Map Err", err)
		t.FailNow()
	}
}

func TestAddRenderMapWithType(t *testing.T) {
	var vm ViewManager

	var i TestRender

	err := vm.AddRenderMapWithType("x", "json", &i)
	if err != "uninit" {
		t.Log("Test View Mananger Add Render Map Err", err)
		t.FailNow()
	}

	err = vm.Init()
	if err != "success" {
		t.Log("Test View Mananger Add Render Map Err", err)
		t.FailNow()
	}

	err = vm.AddRenderMapWithType("x", "json", &i)
	if err != "success" {
		t.Log("Test View Mananger Add Render Map Err", err)
		t.FailNow()
	}

	err = vm.AddRenderMapWithType("", "json", &i)
	if err != "emptystring" {
		t.Log("Test View Mananger Add Render Map Err", err)
		t.FailNow()
	}

	err = vm.AddRenderMapWithType("x", "json1", &i)
	if err != "errortype" {
		t.Log("Test View Mananger Add Render Map Err", err)
		t.FailNow()
	}
}

// 错误码
// success
// regempty 这个字符串没有了对应的 reflect.type
// uninit
func TestGetRender(t *testing.T) {
	var vm ViewManager

	var i TestRender
	err, _ := vm.GetRender("x_json")
	if err != "uninit" {
		t.Log("Test View Mananger Get Render Map Err", err)
		t.FailNow()
	}

	err = vm.Init()
	if err != "success" {
		t.Log("Test View Mananger Get Render Err", err)
		t.FailNow()
	}

	err = vm.AddRenderMap("x_json", &i)
	if err != "success" {
		t.Log("Test View Mananger Get Render Map Err", err)
		t.FailNow()
	}

	errGet, i2 := vm.GetRender("x_json")
	if errGet != "success" && i2.Kind().String() != "Struct" {
		t.Log("Test View Mananger Get Render Map Err", err)
		t.FailNow()
	}
}

// 错误码
// success
// emptystring 名字为空
// errortype 类型错误
// 此接口预留给以后做 fatory 相关的内容
func TestGetRenderByType(t *testing.T) {
	var vm ViewManager

	var i TestRender

	err := vm.Init()
	if err != "success" {
		t.Log("Test View Mananger Get Render By Type Err", err)
		t.FailNow()
	}

	err = vm.AddRenderMap("x.json", &i)
	if err != "success" {
		t.Log("Test View Mananger Get Render Map By Type Err", err)
		t.FailNow()
	}

	err, _ = vm.GetRenderByType("", "json")
	if err != "emptystring" {
		t.Log("Test View Manager Get Render By Type Err", err)
		t.FailNow()
	}

	err, _ = vm.GetRenderByType("x", "json1")
	if err != "errortype" {
		t.Log("Test View Manager Get Render By Type Err", err)
		t.FailNow()
	}

	errGet, i2 := vm.GetRenderByType("x", "json")
	if errGet != "success" && i2.Kind().String() != "Struct" {
		t.Log("Test View Mananger Get Render Map By Type Err", err)
		t.FailNow()
	}
}
