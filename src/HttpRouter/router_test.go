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

func (t *TestControl) Get(w *http.ResponseWriter, r *http.Request) {
}

func (t *TestControl) Put(w *http.ResponseWriter, r *http.Request) {
}

func (t *TestControl) Post(w *http.ResponseWriter, r *http.Request) {
}

func (t *TestControl) Delete(w *http.ResponseWriter, r *http.Request) {
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

func TestResolveURLToRESTData(t *testing.T) {
	var h HttpRouter
	err, szPath, tbParam := h.ResolveURLToRESTData("")
	if err != "needbeign/" || szPath != "" || tbParam != nil {
		t.Log("Resolve URL Err 1")
		t.FailNow()
	}

	err, szPath, tbParam = h.ResolveURLToRESTData("A")
	if err != "needbeign/" || szPath != "" || tbParam != nil {
		t.Log("Resolve URL Err 1.5")
		t.FailNow()
	}

	// /
	err, szPath, tbParam = h.ResolveURLToRESTData("/")
	if err != "success" || szPath != "/" || tbParam != nil {
		t.Log("Resolve URL Err 1.55")
		t.FailNow()
	}

	// /A
	err, szPath, tbParam = h.ResolveURLToRESTData("/A")
	if err != "success" || szPath != "/A" || tbParam != nil {
		t.Log("Resolve URL Err 2", err, szPath, tbParam)
		t.FailNow()
	}

	// /A:n
	err, szPath, tbParam = h.ResolveURLToRESTData("/A:1000")
	if err != "success" || szPath != "/A" || len(tbParam) != 1 || tbParam["A"] != "1000" {
		t.Log("Resolve URL Err 3")
		t.FailNow()
	}

	// /A:n:
	err, szPath, tbParam = h.ResolveURLToRESTData("/A:1000:")
	if err != "errexpr" {
		t.Log("Resolve URL Err 3.1")
		t.FailNow()
	}

	// /A:
	err, szPath, tbParam = h.ResolveURLToRESTData("/A:")
	if err != "errexpr" {
		t.Log("Resolve URL Err 3.2")
		// t.FailNow()
	}

	// /:n
	err, szPath, tbParam = h.ResolveURLToRESTData("/:n")
	if err != "errexpr" {
		t.Log("Resolve URL Err 3.3")
		// t.FailNow()
	}

	// /:
	err, szPath, tbParam = h.ResolveURLToRESTData("/:")
	if err != "errexpr" {
		t.Log("Resolve URL Err 3.4")
		t.FailNow()
	}

	// /A:n/B
	err, szPath, tbParam = h.ResolveURLToRESTData("/A:1000/B")
	if err != "success" || szPath != "/A/B" || len(tbParam) != 1 || tbParam["A"] != "1000" {
		t.Log("Resolve URL Err 4")
		t.FailNow()
	}

	// /A:n/B:m
	err, szPath, tbParam = h.ResolveURLToRESTData("/A:1000/B:2000")
	if err != "success" || szPath != "/A/B" || len(tbParam) != 2 || tbParam["A"] != "1000" || tbParam["B"] != "2000" {
		t.Log("Resolve URL Err 5")
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
