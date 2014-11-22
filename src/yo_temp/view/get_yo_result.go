package view

import (
	"encoding/json"
	"net/http"
)

type YOMsg struct {
	From     string `json:"from"`
	Msg      string `json:"msg"`
	Senddate int    `json:"senddate"`
}

type GetYOData struct {
	Count int     `json:"count"`
	Msgs  []YOMsg `json:"msgs"`
}

// 思考怎么实现数据上的继承
type IGetYO struct {
	Retcode  int       `json:"retcode"`
	Msg      string    `json:"msg"`
	Datetime int       `json:"datetime"`
	Data     GetYOData `json:"data"`
}

type GetYoResult struct{}

func (gy *GetYoResult) Render(i interface{}, w *http.ResponseWriter) {
	lstYO := *(i.(*GetYOData))

	reg := IGetYO{
		Retcode:  200,
		Msg:      "ok",
		Datetime: 10,
		Data:     lstYO,
	}
	encode := json.NewEncoder(*w)
	encode.Encode(reg)
}
