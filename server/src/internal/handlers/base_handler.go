package handlers

import (
	"fmt"
	"net/http"

	"contest-influence/server/internal/handlers/handler_types"
)

type BaseHandler struct {
	Handler http.Handler
	Method  string
}

func (h *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != h.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			switch err := r.(type) {
			case handler_types.HandlerPanic:
				w.WriteHeader(err.StatusCode)
				w.Write([]byte(err.Message))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				_, err = fmt.Fprintf(w, "PANIC: %v\n", err)
				if err != nil {
					w.Write([]byte("Undefined internal error\n"))
				}
			}
		}
	}()

	h.Handler.ServeHTTP(w, r)
}
