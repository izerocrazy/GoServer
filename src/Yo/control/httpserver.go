package yo

import (
	"net/http"
	"reflectmap"
)

type HttpServer struct {
	Server *Server
	Router *ReflectMap
}

/*
函数名：初始化

返回值：error 错误码

success

complaxinit 重复初始化
*/
func (h *HttpServer) Init() string {
	return h.Router.Init()
}

/*
函数名：启动服务

返回值：error 错误码

success

httperr golang http 服务内部错误
*/
func (h *HttpServer) Start(szPort string, req *http.Request) string {

}
