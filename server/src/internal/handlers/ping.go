package handlers

import (
	"net/http"
)

type pingHandler struct {
}

func (h *pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

func NewPingHandler() *BaseHandler {
	return &BaseHandler{
		Handler: &pingHandler{},
		Method:  http.MethodGet,
	}
}
