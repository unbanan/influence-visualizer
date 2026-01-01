package handlers

import (
	"net/http"
)

type PingHandler struct {
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}
