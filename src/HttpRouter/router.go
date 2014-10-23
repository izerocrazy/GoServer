package httprouter

import (
	"base"
	"net/http"
	"reflect"
	"reflectmap"
	"restcontrol"
	"strings"
)

const ViewTypeName = "_view_type"

type HttpRouter struct {
	Map *reflectmap.ReflectMap
}

/*
函数名：初始化 Reflect Map

返回值：error 错误码

success

complaxinit 重复初始化

其余参考 reflectmap.init()
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

其余参考 reflectmap.Add()
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

	tbParam = make(map[string]string)
	if len(szURL) == 1 && szURL[0] == '/' {
		szPath = "/"
		err = "success"
		tbParam[ViewTypeName] = "json"
		return err, szPath, tbParam
	}

	for i, part := range parts {
		if part != "" {
			szPartValue := part
			if i == len(parts)-1 {
				szViewType, szMain := getViewType(part)
				if szViewType == "" {
					tbParam[ViewTypeName] = "json"
				} else {
					tbParam[ViewTypeName] = szViewType
				}
				szPartValue = szMain
			}

			// Base.PrintLog(szPartValue)
			if strings.Contains(szPartValue, ":") {
				tbExpr := strings.Split(szPartValue, ":")
				if len(tbExpr) != 2 || tbExpr[0] == "" || tbExpr[1] == "" {
					return "errexpr", "", nil
				}

				if tbExpr[0] == ViewTypeName {
					return "errexpr", "", nil
				}

				szPath = szPath + "/" + tbExpr[0]
				tbParam[tbExpr[0]] = tbExpr[1]
			} else {
				szPath = szPath + "/" + szPartValue
			}
		}
	}

	if len(tbParam) == 0 {
		tbParam = nil
	}

	return "success", szPath, tbParam
}

func getViewType(szString string) (szType string, szMain string) {
	szType = ""
	szMain = szString
	if strings.Contains(szString, ".") {
		tbExpr := strings.Split(szString, ".")
		if len(tbExpr) == 2 && tbExpr[0] != "" && tbExpr[1] != "" {
			szMain = tbExpr[0]
			szType = tbExpr[1]
		}
	}

	return szType, szMain
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
	errPath, szPath, tbParam := h.ResolveURLToRESTData(r.URL.Path)
	if errPath != "success" {
		Base.PrintErr("HttpRouter ServeHTTP Error: Router Map New a control err: " + errPath + "the path is" + r.URL.Path)
		return
	}

	err, control := h.Map.New(szPath)
	if err != "success" {
		Base.PrintErr("HttpRouter ServeHTTP Error: Router Map New a control err: " + err + "the path is" + r.URL.Path)
		return
	}

	// 需要调用对应 control 的函数
	// Init
	init := control.MethodByName("Init")
	// Base.Fmtprintln(control)
	in2 := make([]reflect.Value, 3)
	in2[0] = reflect.ValueOf(&w)
	in2[1] = reflect.ValueOf(r)
	in2[2] = reflect.ValueOf(tbParam)
	out := init.Call(in2)
	if len(out) != 1 && out[0].String() != "success" {
		Base.PrintErr("HttpRouter ServeHTTP Error: Init one Control Err" + "the path is " + r.URL.Path)
		return
	}

	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(&w)
	in[1] = reflect.ValueOf(r)
	// Get and Post
	if r.Method == "GET" {
		method := control.MethodByName("Get")
		method.Call(in)
	} else if r.Method == "POST" {
		method := control.MethodByName("Post")
		method.Call(in)
	} else if r.Method == "PUT" {
		method := control.MethodByName("Put")
		method.Call(in)
	} else if r.Method == "DELETE" {
		method := control.MethodByName("Delete")
		method.Call(in)
	} else {
		Base.PrintErr("Http Router ServeHTTP Error: Router Map Get A Error Method" + r.Method)
	}
}
