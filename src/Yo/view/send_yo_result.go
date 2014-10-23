package view

import (
	"encoding/json"
	"net/http"
)

type ISendYO struct {
	Retcode  int    `json:"retcode"`
	Msg      string `json:"msg"`
	Datetime int    `json:"datetime"`
}

type SendYoResult struct{}

func (sy *SendYoResult) Render(i interface{}, w *http.ResponseWriter) {
	reg := ISendYO{
		Retcode:  200,
		Msg:      "ok",
		Datetime: 10,
	}
	encode := json.NewEncoder(*w)
	encode.Encode(reg)
}
