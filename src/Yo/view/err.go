package view

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	RetCode   int    `json:"retcode"`
	Msg       string `json:"msg"`
	ErrorCode string `json:"errorcode"`
}

type ErrResult struct{}

func MakeError(err string) (reg interface{}) {
	reg = Error{
		RetCode:   200,
		Msg:       "error",
		ErrorCode: err,
	}

	return reg
}

func (er *ErrResult) Render(i interface{}, w *http.ResponseWriter) {
	szErr := i.(string)
	reg := MakeError(szErr)
	encode := json.NewEncoder(*w)
	encode.Encode(reg)
}
