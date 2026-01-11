package handlers

import (
	"fmt"
	"net/http"

	"contest-influence/server/internal/handlers/handler_types"
)

type baseHandler struct {
	Handler http.Handler
}

func WrapHandler(handler http.Handler) *baseHandler {
	return &baseHandler{Handler: handler}
}

func (h *baseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
