package unit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"contest-influence/server/internal/handlers"
	"contest-influence/server/internal/simulation_types"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimulationGetHandler_Success(t *testing.T) {
	testUUID := uuid.New()
	mock := &InfluenceDBRepoMock{}
	handler := handlers.NewGetSimulationHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/simulation?id="+testUUID.String(), nil)

	handler.ServeHTTP(w, r)

	ExpectStatusCodesEqual(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var result simulation_types.Simulation
	err := json.NewDecoder(w.Body).Decode(&result)
	require.NoError(t, err)
}

func TestSimulationGetHandler_InvalidUUID(t *testing.T) {
	mock := &InfluenceDBRepoMock{}
	handler := handlers.NewGetSimulationHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/simulation?id=invalid-uuid", nil)

	handler.ServeHTTP(w, r)

	ExpectStatusCodesEqual(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func TestSimulationGetHandler_MissingID(t *testing.T) {
	mock := &InfluenceDBRepoMock{}
	handler := handlers.NewGetSimulationHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/simulation", nil)

	handler.ServeHTTP(w, r)

	ExpectStatusCodesEqual(t, http.StatusBadRequest, w.Result().StatusCode)
}
