package handlers

import (
	"net/http"
	"regexp"

	"contest-influence/server/internal/database/influence"
	"contest-influence/server/internal/handlers/params"
)

type RegisterHandlerImpl interface {
	Register(id int64, name string)
}

type registerHandler struct {
	Impl  RegisterHandlerImpl
	Regex *regexp.Regexp
}

func (h *registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

type registerHandlerImpl struct {
	InfluenceDBRepo influence.InfluenceDBRepo
}

func (h *registerHandlerImpl) Register(id int64, name string) {
	err := h.InfluenceDBRepo.Register(id, name)

	if err != nil {
		panic(err)
	}
}

func NewRegisterHandlerImpl(influencedb influence.InfluenceDBRepo) RegisterHandlerImpl {
	return &registerHandlerImpl{
		InfluenceDBRepo: influencedb,
	}
}

func NewRegisterHandler(regex *regexp.Regexp, influencedb influence.InfluenceDBRepo) http.Handler {
	return WrapHandler(
		&registerHandler{
			Impl:  NewRegisterHandlerImpl(influencedb),
			Regex: regex,
		},
	)
}
