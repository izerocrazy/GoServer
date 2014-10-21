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

	err = vm.AddRenderMap("x_json", &i)
	if err != "isexist" {
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

	err, _ = vm.GetRender("x_json")
	if err != "regempty" {
		t.Log("Test View Mananger Get Render Map Err", err)
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
