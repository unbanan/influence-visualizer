package unit

import (
	"errors"
	"testing"

	"contest-influence/server/internal/handlers"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHandlerImpl_Success(t *testing.T) {
	repoMock := &InfluenceDBRepoMock{}
	impl := handlers.NewRegisterHandlerImpl(repoMock)

	impl.Register(42, "username")

	assert.Equal(t, int64(1), repoMock.RegisterCalledCount)
	assert.Equal(t, int64(42), repoMock.LastID)
	assert.Equal(t, "username", repoMock.LastName)
}

func TestRegisterHandlerImpl_RepositoryError(t *testing.T) {
	repoMock := &InfluenceDBRepoMock{
		ShouldReturnError: true,
		ErrorToReturn:     errors.New("duplicate key violation"),
	}
	impl := handlers.NewRegisterHandlerImpl(repoMock)

	assert.Panics(t, func() {
		impl.Register(42, "username")
	})
	assert.Equal(t, int64(1), repoMock.RegisterCalledCount)
}
