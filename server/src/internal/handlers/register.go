package handlers

import (
	"net/http"

	"contest-influence/server/internal/context"
	"contest-influence/server/internal/handlers/params"
)

type IRegisterHandlerImpl interface {
	Register(id int64, name string)
}

type RegisterHandler struct {
	Impl IRegisterHandlerImpl
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := params.GetInt(r, "id")
	name := params.GetString(r, "name")

	h.Impl.Register(id, name)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfuly registered"))
}

func NewRegisterHandler(context *context.Context) *BaseHandler {
	return &BaseHandler{
		Handler: &RegisterHandler{
			Impl: &RegisterHandlerImpl{
				Context: context,
			},
		},
		Method: http.MethodPost,
	}
}

type RegisterHandlerImpl struct {
	Context *context.Context
}

func (h *RegisterHandlerImpl) Register(id int64, name string) {
	repo := h.Context.InfluenceDBRepo

	repo.Register(id, name)
}
