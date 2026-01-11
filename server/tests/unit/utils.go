package unit

import (
	"net/http"
	"regexp"
	"testing"

	"contest-influence/server/internal/database/influence"
	"contest-influence/server/internal/handlers"
	"contest-influence/server/internal/simulation_types"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func ExpectStatusCodesEqual(t *testing.T, expected, actual int) {
	assert.Equalf(t, expected, actual, "Wrong status code")
}

type InfluenceDBRepoMock struct {
	RegisterCalledCount int64
	LastID              int64
	LastName            string
	ShouldReturnError   bool
	ErrorToReturn       error
}

func (m *InfluenceDBRepoMock) Register(id int64, name string) error {
	m.RegisterCalledCount++
	m.LastID = id
	m.LastName = name
	if m.ShouldReturnError {
		return m.ErrorToReturn
	}
	return nil
}

func (m *InfluenceDBRepoMock) GetSimulationPlayers(uuid.UUID) ([]influence.User, error) {
	return nil, nil
}

func (m *InfluenceDBRepoMock) GetSimulation(uuid.UUID) (*simulation_types.Simulation, error) {
	return nil, nil
}

func (m *InfluenceDBRepoMock) Close() error {
	return nil
}

func NewTestRegisterHandler(repo influence.InfluenceDBRepo) http.Handler {
	return handlers.NewRegisterHandler(
		regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`),
		repo,
	)
}
