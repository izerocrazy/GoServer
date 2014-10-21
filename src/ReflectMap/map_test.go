package reflectmap

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	var r ReflectMap
	err := r.Init()
	if err != "success" {
		t.Log("init err", err)
		t.FailNow()
	}

	err = r.Init()
	if err != "complaxinit" {
		t.Log("init err", err)
		t.FailNow()
	}
}

func TestAdd(t *testing.T) {
	var r ReflectMap
	var x int
	err := r.Add("int", x)
	if err != "uninit" {
		t.Log("Add Err", err)
		t.FailNow()
	}

	err = r.Init()
	if err != "success" {
		t.Log("Add Err", err)
		t.FailNow()
	}

	err = r.Add("int", x)
	if err != "success" {
		t.Log("add int err, ", err)
		t.FailNow()
	}

	err = r.Add("int", x)
	if err != "isexist" {
		t.Log("second add int err, ", err)
		t.FailNow()
	}
}

type ForTest struct {
	I int
}

func (f *ForTest) Init(x int) int {
	f.I = x
	return f.I
}

func TestNew(t *testing.T) {
	var r ReflectMap
	var x ForTest
	err := r.Add("int", x)
	if err != "uninit" {
		t.Log("Add Err", err)
		t.FailNow()
	}

	err = r.Init()
	if err != "success" {
		t.Log("add int err, ", err)
		t.FailNow()
	}

	err = r.Add("int", x)
	if err != "success" {
		t.Log("add int err, ", err)
		t.FailNow()
	}

	err2, control := r.New("int")
	if err2 != "success" {
		t.Log("New Int err, ", err)
		t.FailNow()
	}

	err3, _ := r.New("int2")
	if err3 != "regempty" {
		t.Log("New Int2 err", err)
		t.FailNow()
	}

	init := control.MethodByName("Init")
	in := make([]reflect.Value, 1)
	// ct := &Context{ResponseWriter: w, Request: r, Params: params}
	ct := 10
	in[0] = reflect.ValueOf(ct)
	// in[1] = reflect.ValueOf("int")
	out := init.Call(in)
	if len(out) != 1 || out[0].Int() != 10 {
		t.Log("New Int3 err call func fail")
		t.FailNow()
	}
}

func TestNewInterface(t *testing.T) {
	var r ReflectMap
	r.Init()

	var x int
	x = 1
	err := r.Add("int", x)
	if err != "success" {
		t.Log("add int err, ", err)
		t.FailNow()
	}

	err2, _ := r.NewInterface("int")
	if err2 != "success" {
		t.Log("New Int err, ", err)
		t.FailNow()
	}

	err3, _ := r.NewInterface("int2")
	if err3 != "regempty" {
		t.Log("New Int2 err", err)
		t.FailNow()
	}

}
