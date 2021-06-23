package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApiHandler struct {
	Method      string
	Uri         string
	HandlerFunc http.HandlerFunc
}

func SetUpHandlers(handlers ...ApiHandler) http.Handler {
	r := chi.NewRouter()

	for _, handler := range handlers {
		r.MethodFunc(handler.Method, handler.Uri, handler.HandlerFunc)
	}

	return r
}
