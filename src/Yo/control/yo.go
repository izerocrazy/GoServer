package view

import (
	"base"
	"net/http"
	"yo/module"
)

type YoControl struct {
	BaseControl
}

// 获取未读的 yo
func (yc *YoControl) Get(w *http.ResponseWriter, r *http.Request) {
}

func (yc *YoControl) Post(w *http.ResponseWriter, r *http.Request) {
}
