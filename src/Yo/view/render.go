package view

import (
	"net/http"
	"yo/module"
)

type Render interface {
	Render(err string, user *module.UserData, w *http.ResponseWriter)
}
