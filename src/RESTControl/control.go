package restcontrol

import (
	"net/http"
)

type RESTControl interface {
	Init(w *http.ResponseWriter, r *http.Request)

	Get()

	Put()

	Post()

	Delete()
}
