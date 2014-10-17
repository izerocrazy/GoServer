package httprouter

import (
	"base"
	"net/http"
	"reflectmap"
	"restcontrol"
)

type HttpRouter struct {
	Map *reflectmap.ReflectMap
}

/*
函数名：初始化 Reflect Map

返回值：error 错误码

success

complaxinit 重复初始化
*/
func (h *HttpRouter) Init() string {
	var err string
	if h.Map == nil {
		h.Map = new(reflectmap.ReflectMap)
		err = h.Map.Init()
	} else {
		err = "complaxinit"
	}

	return err
}

/*
函数名：增加路径对应的响应类型

返回值：error 错误码

success

uninit 未初始化

isexist 这个字符串已经有了对应的控制器
*/
func (h *HttpRouter) AddControl(szPath string, control restcontrol.RESTControl) string {
	var err string
	if h.Map == nil {
		err = "uninit"
	} else {
		err = h.Map.Add(szPath, control)
	}

	return err
}

/*
函数名：实现 golang http 库中的 Handle 接口，完成路由器
// 错误有：
// 对象错误：1、map 未初始化
// 参数错误：2、ResponseWriter 的错误；3、Request 的错误
*/
func (h *HttpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Map == nil {
		Base.PrintErr("HttpRouter ServeHTTP Error: HttpRouter.ReflectMap is un Init")
	}

	// 从 map 中取出一个对象
	err, _ := h.Map.New(r.URL.Path)
	if err != "success" {
		Base.PrintErr("HttpRouter ServeHTTP Error: Map New a control err: " + err)
	}

	// 需要调用对应 control 的函数
}
