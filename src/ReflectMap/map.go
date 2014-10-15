package reflectmap

import (
	"reflect"
)

type ReflectMap struct {
	Table map[string]reflect.Type
}

/*
函数名：初始化 Reflect Map

返回值：error 错误码

success

complaxinit 重复初始化
*/
func (r *ReflectMap) Init() string {
	if r.Table == nil {
		r.Table = make(map[string]reflect.Type)
		return "success"
	}

	return "complaxinit"
}

/*
函数名：注册 string 和一个 reflect.type 的关系

参数：szRegName 待注册的名字, rt reflect.Type

返回值：error 错误码

success

isexist 这个字符串已经有了对应的 reflect.type
*/
func (r *ReflectMap) Add(szRegName string, i interface{}) string {
	if r.Table[szRegName] != nil {
		return "isexist"
	}

	rt := reflect.ValueOf(i)
	r.Table[szRegName] = rt.Type()
	return "success"
}

/*
函数名：传入 string，生成一个类型对应的指针

参数：szRegName 注册名称

返回值：error 错误码

success

regempty 这个字符串已经有了对应的 reflect.type
*/
func (r *ReflectMap) New(szRegName string) (err string, i interface{}) {
	t := r.Table[szRegName]
	if t != nil {
		i = reflect.New(t).Interface()
		err = "success"
	} else {
		err = "regempty"
	}

	return err, i
}
