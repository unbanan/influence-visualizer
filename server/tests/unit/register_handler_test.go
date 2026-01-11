package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler_ValidRequest(t *testing.T) {
	mock := &InfluenceDBRepoMock{}
	handler := NewTestRegisterHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=123&name=testuser", nil)

	handler.ServeHTTP(w, r)

	assert.Equal(t, int64(1), mock.RegisterCalledCount)
	assert.Equal(t, int64(123), mock.LastID)
	assert.Equal(t, "testuser", mock.LastName)
	ExpectStatusCodesEqual(t, http.StatusOK, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "Successfuly registered")
}

func TestRegisterHandler_InvalidName(t *testing.T) {
	testCases := []struct {
		name     string
		queryStr string
	}{
		{"too short", "?id=1&name=ab"},
		{"too long", "?id=1&name=verylongnamethatexceedslimit"},
		{"special chars", "?id=1&name=user@name"},
		{"spaces", "?id=1&name=user%20name"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &InfluenceDBRepoMock{}
			handler := NewTestRegisterHandler(mock)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/api/v1/register"+tc.queryStr, nil)

			handler.ServeHTTP(w, r)

			assert.Equal(t, int64(0), mock.RegisterCalledCount)
			ExpectStatusCodesEqual(t, http.StatusBadRequest, w.Result().StatusCode)
			assert.Contains(t, w.Body.String(), "Invalid name")
		})
	}
}

func TestRegisterHandler_MissingParams(t *testing.T) {
	testCases := []struct {
		name     string
		queryStr string
	}{
		{"missing id", "?name=testuser"},
		{"missing name", "?id=123"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &InfluenceDBRepoMock{}
			handler := NewTestRegisterHandler(mock)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/api/v1/register"+tc.queryStr, nil)

			handler.ServeHTTP(w, r)

			assert.Equal(t, int64(0), mock.RegisterCalledCount)
			ExpectStatusCodesEqual(t, http.StatusBadRequest, w.Result().StatusCode)
		})
	}
}

func TestRegisterHandler_InvalidID(t *testing.T) {
	mock := &InfluenceDBRepoMock{}
	handler := NewTestRegisterHandler(mock)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=invalid&name=testuser", nil)

	handler.ServeHTTP(w, r)

	assert.Equal(t, int64(0), mock.RegisterCalledCount)
	ExpectStatusCodesEqual(t, http.StatusBadRequest, w.Result().StatusCode)
	assert.Contains(t, w.Body.String(), "field id has wrong format")
}
