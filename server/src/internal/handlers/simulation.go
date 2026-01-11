package handlers

import (
	"encoding/json"
	"net/http"

	"contest-influence/server/internal/database/influence"
	"contest-influence/server/internal/handlers/params"
	"contest-influence/server/internal/simulation_types"

	"github.com/google/uuid"
)

type SimulationGetHandlerImpl interface {
	GetSimulation(uuid.UUID) *simulation_types.Simulation
}

type simulationGetHandler struct {
	Impl SimulationGetHandlerImpl
}

func (h *simulationGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := params.GetUUID(r, "id")
	simulation := h.Impl.GetSimulation(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&simulation)
}

type simulationGetHandlerImpl struct {
	InfluenceDBRepo influence.InfluenceDBRepo
}

func (h *simulationGetHandlerImpl) GetSimulation(id uuid.UUID) *simulation_types.Simulation {
	sim, err := h.InfluenceDBRepo.GetSimulation(id)
	if err != nil {
		panic(err)
	}
	return sim
}

func NewSimulationGetHandlerImpl(influencedb influence.InfluenceDBRepo) SimulationGetHandlerImpl {
	return &simulationGetHandlerImpl{InfluenceDBRepo: influencedb}
}

func NewGetSimulationHandler(influencedb influence.InfluenceDBRepo) http.Handler {
	return WrapHandler(
		&simulationGetHandler{
			Impl: NewSimulationGetHandlerImpl(influencedb),
		},
	)
}
