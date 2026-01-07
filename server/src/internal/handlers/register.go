package handlers

import (
	"net/http"
	"regexp"

	"contest-influence/server/internal/handlers/params"
	"contest-influence/server/internal/repos"
)

type IRegisterHandlerImpl interface {
	Register(id int64, name string)
}

type RegisterHandler struct {
	Impl  IRegisterHandlerImpl
	Regex *regexp.Regexp
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := params.GetInt(r, "id")
	name := params.GetString(r, "name")

	if !h.Regex.MatchString(name) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid name"))
		return
	}

	h.Impl.Register(id, name)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfuly registered"))
}

func NewRegisterHandler(regex *regexp.Regexp, influencedb repos.InfluenceDBRepo) *BaseHandler {
	return &BaseHandler{
		Handler: &RegisterHandler{
			Impl: &RegisterHandlerImpl{
				InfluenceDBRepo: influencedb,
			},
			Regex: regex,
		},
		Method: http.MethodPost,
	}
}

type RegisterHandlerImpl struct {
	InfluenceDBRepo repos.InfluenceDBRepo
}

func (h *RegisterHandlerImpl) Register(id int64, name string) {
	err := h.InfluenceDBRepo.Register(id, name)

	if err != nil {
		panic(err.Error())
	}
}
