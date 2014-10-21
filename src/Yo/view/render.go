package view

import (
	"net/http"
)

type Render interface {
	Render(i interface{}, w *http.ResponseWriter)
}
