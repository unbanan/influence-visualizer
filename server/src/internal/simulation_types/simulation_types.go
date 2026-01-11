package simulation_types

import (
	"time"

	proto "contest-influence/proto/simulation"
)

type Map struct {
	Data proto.Map `json:"map"`
	Name string    `json:"name"`
}

type Simulation struct {
	Map        Map              `json:"map"`
	Data       proto.Simulation `json:"data"`
	Players    []string         `json:"players"`
	QueuedAt   time.Time        `json:"queued_at"`
	StartedAt  time.Time        `json:"starterd_at"`
	FinishedAt time.Time        `json:"finished_at"`
	State      string           `json:"state"`
}

func NewSimulation() *Simulation {
	return &Simulation{
		Players: make([]string, 0),
	}
}
