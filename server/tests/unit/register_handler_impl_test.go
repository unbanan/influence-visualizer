package unit

import (
	"errors"
	"testing"

	"contest-influence/server/internal/handlers"
	"contest-influence/server/internal/repos"

	"github.com/stretchr/testify/assert"
)

type InfluenceDBRepoMock struct {
	repos.InfluenceDBRepo
	RegisterCalledCount int64
	LastID              int64
	LastName            string
	ShouldReturnError   bool
	ErrorMessage        string
}

func (m *InfluenceDBRepoMock) Register(id int64, name string) error {
	m.RegisterCalledCount++
	m.LastID = id
	m.LastName = name
	if m.ShouldReturnError {
		return errors.New(m.ErrorMessage)
	}
	return nil
}

func TestRegisterHandlerImpl_Success(t *testing.T) {
	repoMock := &InfluenceDBRepoMock{}
	impl := &handlers.RegisterHandlerImpl{
		InfluenceDBRepo: repoMock,
	}

	impl.Register(42, "username")

	assert.Equal(t, int64(1), repoMock.RegisterCalledCount)
	assert.Equal(t, int64(42), repoMock.LastID)
	assert.Equal(t, "username", repoMock.LastName)
}

func TestRegisterHandlerImpl_RepositoryError(t *testing.T) {
	repoMock := &InfluenceDBRepoMock{
		ShouldReturnError: true,
		ErrorMessage:      "duplicate key violation",
	}
	impl := &handlers.RegisterHandlerImpl{
		InfluenceDBRepo: repoMock,
	}

	assert.PanicsWithValue(t, "duplicate key violation", func() {
		impl.Register(42, "username")
	})
	assert.Equal(t, int64(1), repoMock.RegisterCalledCount)
}
