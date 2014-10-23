package restcontrol

import (
	"net/http"
)

type RESTControl interface {
	Init(w *http.ResponseWriter, r *http.Request, tbParam map[string]string) (err string)

	Get(w *http.ResponseWriter, r *http.Request)

	Put(w *http.ResponseWriter, r *http.Request)

	Post(w *http.ResponseWriter, r *http.Request)

	Delete(w *http.ResponseWriter, r *http.Request)
}
