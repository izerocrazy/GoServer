package httprouter

import (
	"net/http"
	"testing"
)

// 这是因为需要初始化 map
// 它的错误就是重复初始化
func TestInit(t *testing.T) {
	var h HttpRouter

	err := h.Init()
	if err != "success" {
		t.Log("Init Error:", err)
		t.FailNow()
	}

	err = h.Init()
	if err != "complaxinit" {
		t.Log("Init Error:", err)
		t.FailNow()
	}
}

type TestControl struct {
}

func (t *TestControl) Init(w *http.ResponseWriter, r *http.Request) {
}

func (t *TestControl) Get() {
}

func (t *TestControl) Put() {
}

func (t *TestControl) Post() {
}

func (t *TestControl) Delete() {
}

// 对象错误：1、map 未初始化
// 参数错误：1、map 中已有此路径
func TestAddControl(t *testing.T) {
	var h HttpRouter
	szHandle := "/"
	handler := new(TestControl)

	err := h.AddControl(szHandle, handler)
	if err != "uninit" {
		t.Log("AddControl Error:", err, "need uninit")
		t.FailNow()
	}

	err = h.Init()
	if err != "success" {
		t.Log("AddControl Error: Init router error", err)
		t.FailNow()
	}

	err = h.AddControl(szHandle, handler)
	if err != "success" {
		t.Log("Add Control Error:", err, "need success")
		t.FailNow()
	}

	err = h.AddControl(szHandle, handler)
	if err != "isexist" {
		t.Log("Add Control Error", err, "need isexist")
		t.FailNow()
	}
}

// 因为原生的 DefaultServeMux 对外只有这一个接口，所以是不是也能做到？
// func TestServeHTTP(t *testing.T) {
// 	var h HttpRouter
// 	var w http.ResponseWriter
// 	var r http.Request
// 	err := h.ServeHTTP(w, &r)
// 	if err != "uninit" {
// 		t.Log("ServeHTTP Error:", err, "need uninit")
// 		t.FailNow()
// 	}

// 	err = h.Init()
// 	if err != "success" {
// 		t.Log("ServeHTTP Error: init router error", err)
// 		t.FailNow()
// 	}

// 	err = h.ServeHTTP(w, &r)
// 	if err != "mapempty" {
// 		t.Log("ServeHTTP Error:", err, "need mapempty")
// 		t.FailNow()
// 	}

// 	szHandle := "/"
// 	var handler int // 暂用 int 代替
// 	err = h.AddControl(szHandle, handler)
// 	if err != "success" {
// 		t.Log("ServeHTTP Error: add control error", err)
// 		t.FailNow()
// 	}

// 	// 缺参数错误

// 	err = h.ServeHTTP(w, &r)
// 	if err != "success" {
// 		t.Log("ServeHTTP Error:", err, "need success")
// 		t.FailNow()
// 	}
// }
