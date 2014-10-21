package view

import ()

type Error struct {
	RetCode   int    `json:"retcode"`
	Msg       string `json:"msg"`
	ErrorCode string `json:"errorcode"`
}

func MakeError(err string) (reg interface{}) {
	reg = Error{
		RetCode:   200,
		Msg:       "error",
		ErrorCode: err,
	}

	return reg
}
