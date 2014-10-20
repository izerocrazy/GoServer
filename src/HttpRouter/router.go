package httprouter

import (
	"base"
	"net/http"
	"reflect"
	"reflectmap"
	"restcontrol"
	"strings"
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

func (h *HttpRouter) ResolveURLToRESTData(szURL string) (err string, szPath string, tbParam map[string]string) {
	szPath = ""
	tbParam = nil

	parts := strings.Split(szURL, "/")
	// 第一个必须是 ""
	if len(szURL) < 1 || szURL[0] != '/' || parts[0] != "" {
		return "needbeign/", szPath, tbParam
	}

	if len(szURL) == 1 {
		szPath = "/"
		err = "success"
		return err, szPath, tbParam
	}

	tbParam = make(map[string]string)
	for _, part := range parts {
		if part != "" {
			if strings.Contains(part, ":") {
				tbExpr := strings.Split(part, ":")
				if len(tbExpr) == 2 && tbExpr[0] != "" && tbExpr[1] != "" {
					szPath = szPath + "/" + tbExpr[0]
					tbParam[tbExpr[0]] = tbExpr[1]
				} else {
					// fmt.Println(part, len(tbExpr), tbExpr[0])
					return "errexpr", "", nil
				}
			} else {
				szPath = szPath + "/" + part
			}
		}
	}

	if len(tbParam) == 0 {
		tbParam = nil
	}

	return "success", szPath, tbParam
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
		return
	}

	// 从 map 中取出一个对象:New
	err, control := h.Map.New(r.URL.Path)
	if err != "success" {
		Base.PrintErr("HttpRouter ServeHTTP Error: Router Map New a control err: " + err + "the path is" + r.URL.Path)
		return
	}

	// 需要调用对应 control 的函数
	// Init
	init := control.MethodByName("Init")
	Base.Fmtprintln(control)
	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(&w)
	in[1] = reflect.ValueOf(r)
	init.Call(in)

	// Get and Post
	if r.Method == "Get" {
		method := control.MethodByName("Get")
		method.Call(in)
	} else if r.Method == "Post" {
		method := control.MethodByName("Post")
		method.Call(in)
	} else if r.Method == "put" {
		method := control.MethodByName("Put")
		method.Call(in)
	} else if r.Method == "Delete" {
		method := control.MethodByName("Delete")
		method.Call(in)
	}
}
