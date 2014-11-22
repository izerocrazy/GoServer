package yo

import (
	"net/http"
	"testing"
	"yo/module"
)

func TestInit(t *testing.T) {
	var h HttpServer
	err := h.Init()
	if err != "success" {
		t.Log("HttpServer Init Err", err)
		t.FailNow()
	}

	err = h.Init()
	if err != "complaxinit" {
		t.Log("HttpServer Init Err", err)
		t.FailNow()
	}
}

type TestControl struct {
}

func (t *TestControl) Init(w *http.ResponseWriter, r *http.Request, tbParam map[string]string) string {
	return "success"
}

func (t *TestControl) Get(w *http.ResponseWriter, r *http.Request) {
}

func (t *TestControl) Put(w *http.ResponseWriter, r *http.Request) {
}

func (t *TestControl) Post(w *http.ResponseWriter, r *http.Request) {
}

func (t *TestControl) Delete(w *http.ResponseWriter, r *http.Request) {
}

// 错误有：1、未初始化；2、重复绑定
func TestAddControl(t *testing.T) {
	var h HttpServer

	handle := new(TestControl)
	err := h.AddControl("/", handle)
	if err != "uninit" {
		t.Log("HttpServer AddControl Error:", err)
		t.FailNow()
	}

	err = h.Init()
	if err != "success" {
		t.Log("HttpServer AddControl Error: init error")
		t.FailNow()
	}

	err = h.AddControl("/", handle)
	if err != "success" {
		t.Log("HttpServer AddControl Error:", err)
		t.FailNow()
	}

	err = h.AddControl("/", handle)
	if err != "isexist" {
		t.Log("HttpServer AddControl Error:", err)
		t.FailNow()
	}
}

type TestRender struct{}

func (tr *TestRender) Render(err string, user *module.UserData, w *http.ResponseWriter) {

}

// 错误的情况有：
// 0、未初始化
// ======================
// 2、接口已经使用
// 3、已经启用
// func TestStart(t *testing.T) {
// 	var h HttpServer
// 	err := h.Start(":8080")
// 	if err != "uninit" {
// 		t.Log("HttpServer Start Err", err)
// 		t.FailNow()
// 	}

// 	err = h.Init()
// 	if err != "success" {
// 		t.Log("HttpServer Start Err", err)
// 		t.FailNow()
// 	}

// 	err = h.Start(":8080")
// 	if err != "success" {
// 		t.Log("HttpServer Start Err", err)
// 		t.FailNow()
// 	}

// 	// 确认这样可以测？
// 	err = h.Start(":8080")
// 	if err == "success" || err == "init" {
// 		t.Log("HttpServer Start Err", err)
// 		t.FailNow()
// 	}
// }
