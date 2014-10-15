package reflectmap

import (
	// "fmt"
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
	r.Init()

	var x int
	err := r.Add("int", x)
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

func TestNew(t *testing.T) {
	var r ReflectMap
	r.Init()

	var x int
	x = 1
	err := r.Add("int", x)
	if err != "success" {
		t.Log("add int err, ", err)
		t.FailNow()
	}

	err2, _ := r.New("int")
	if err2 != "success" {
		t.Log("New Int err, ", err)
		t.FailNow()
	}

	err3, _ := r.New("int2")
	if err3 != "regempty" {
		t.Log("New Int2 err", err)
		t.FailNow()
	}
}
