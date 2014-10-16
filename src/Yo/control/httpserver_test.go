package yo

import (
	"testing"
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

// 错误有：1、未初始化；2、重复绑定
func TestAddHandle(t *testing.T) {

}

// 错误的情况有：
// 0、未初始化
// 1、未绑定接口服务
// 2、接口已经使用
// 3、已经启用
func TestStart(t *testing.T) {
	var h HttpServer
	err := h.Start(":8080", nil)
	if err != "success" {
		t.Log("HttpServer Start Err", err)
	}
}
