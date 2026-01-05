package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"contest-influence/server/internal/handlers"

	"github.com/stretchr/testify/assert"
)

// MockRegisterHandlerImpl is a mock implementation for testing
type MockRegisterHandlerImpl struct {
	RegisterCalled bool
	LastID         int64
	LastName       string
}

func (m *MockRegisterHandlerImpl) Register(id int64, name string) {
	m.RegisterCalled = true
	m.LastID = id
	m.LastName = name
}

func TestRegisterHandler_Success(t *testing.T) {
	mockImpl := &MockRegisterHandlerImpl{}
	handler := &handlers.RegisterHandler{
		Impl: mockImpl,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=123&name=testuser", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	ExpectStatusCodesEqual(t, http.StatusOK, w.Code)
	assert.True(t, mockImpl.RegisterCalled, "Expected Register to be called")
	assert.Equal(t, int64(123), mockImpl.LastID, "Expected ID to be 123")
	assert.Equal(t, "testuser", mockImpl.LastName, "Expected name to be 'testuser'")
	assert.Equal(t, "Successfuly registered", w.Body.String(), "Expected success message")
}

func TestRegisterHandler_MissingIDParameter(t *testing.T) {
	mockImpl := &MockRegisterHandlerImpl{}
	handler := &handlers.RegisterHandler{
		Impl: mockImpl,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register?name=testuser", nil)
	w := httptest.NewRecorder()

	assert.Panics(t, func() {
		handler.ServeHTTP(w, req)
	}, "Expected panic for missing id parameter")
}

func TestRegisterHandler_MissingNameParameter(t *testing.T) {
	mockImpl := &MockRegisterHandlerImpl{}
	handler := &handlers.RegisterHandler{
		Impl: mockImpl,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=123", nil)
	w := httptest.NewRecorder()

	assert.Panics(t, func() {
		handler.ServeHTTP(w, req)
	}, "Expected panic for missing name parameter")
}

func TestRegisterHandler_InvalidIDFormat(t *testing.T) {
	mockImpl := &MockRegisterHandlerImpl{}
	handler := &handlers.RegisterHandler{
		Impl: mockImpl,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=invalid&name=testuser", nil)
	w := httptest.NewRecorder()

	assert.Panics(t, func() {
		handler.ServeHTTP(w, req)
	}, "Expected panic for invalid id format")
}

func TestRegisterHandler_EmptyName(t *testing.T) {
	mockImpl := &MockRegisterHandlerImpl{}
	handler := &handlers.RegisterHandler{
		Impl: mockImpl,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=123&name=", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	ExpectStatusCodesEqual(t, http.StatusOK, w.Code)
	assert.Equal(t, "", mockImpl.LastName, "Expected empty name")
}

func TestRegisterHandler_NegativeID(t *testing.T) {
	mockImpl := &MockRegisterHandlerImpl{}
	handler := &handlers.RegisterHandler{
		Impl: mockImpl,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=-1&name=testuser", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	ExpectStatusCodesEqual(t, http.StatusOK, w.Code)
	assert.Equal(t, int64(-1), mockImpl.LastID, "Expected ID to be -1")
}

func TestRegisterHandler_LargeID(t *testing.T) {
	mockImpl := &MockRegisterHandlerImpl{}
	handler := &handlers.RegisterHandler{
		Impl: mockImpl,
	}

	largeID := int64(9223372036854775807) // max int64
	req := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=9223372036854775807&name=testuser", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	ExpectStatusCodesEqual(t, http.StatusOK, w.Code)
	assert.Equal(t, largeID, mockImpl.LastID, "Expected ID to be max int64")
}

func TestRegisterHandler_SpecialCharactersInName(t *testing.T) {
	mockImpl := &MockRegisterHandlerImpl{}
	handler := &handlers.RegisterHandler{
		Impl: mockImpl,
	}

	specialName := "test@user#123"
	req := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=123&name="+specialName, nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	ExpectStatusCodesEqual(t, http.StatusOK, w.Code)
	assert.Equal(t, specialName, mockImpl.LastName, "Expected name with special characters")
}

func TestRegisterHandler_VariousInputs(t *testing.T) {
	testCases := []struct {
		name         string
		id           string
		userName     string
		expectedID   int64
		expectedName string
	}{
		{
			name:         "Zero ID",
			id:           "0",
			userName:     "zero_user",
			expectedID:   0,
			expectedName: "zero_user",
		},
		{
			name:         "URL encoded name",
			id:           "123",
			userName:     "John+Doe",
			expectedID:   123,
			expectedName: "John Doe",
		},
		{
			name:         "Unicode characters",
			id:           "456",
			userName:     "Пользователь",
			expectedID:   456,
			expectedName: "Пользователь",
		},
		{
			name:         "Very long name",
			id:           "789",
			userName:     "verylongnamethatexceedsnormallimitsbutshouldbetested",
			expectedID:   789,
			expectedName: "verylongnamethatexceedsnormallimitsbutshouldbetested",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockImpl := &MockRegisterHandlerImpl{}
			handler := &handlers.RegisterHandler{
				Impl: mockImpl,
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/register?id="+tc.id+"&name="+tc.userName, nil)
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)

			ExpectStatusCodesEqual(t, http.StatusOK, w.Code)
			assert.True(t, mockImpl.RegisterCalled, "Expected Register to be called")
			assert.Equal(t, tc.expectedID, mockImpl.LastID, "Expected ID to match")
			assert.Equal(t, tc.expectedName, mockImpl.LastName, "Expected name to match")
		})
	}
}

func TestRegisterHandler_InvalidInputs(t *testing.T) {
	testCases := []struct {
		name string
		url  string
	}{
		{
			name: "ID as float",
			url:  "/api/v1/register?id=123.45&name=testuser",
		},
		{
			name: "ID overflow",
			url:  "/api/v1/register?id=99999999999999999999&name=testuser",
		},
		{
			name: "ID as letters",
			url:  "/api/v1/register?id=abc&name=testuser",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockImpl := &MockRegisterHandlerImpl{}
			handler := &handlers.RegisterHandler{
				Impl: mockImpl,
			}

			req := httptest.NewRequest(http.MethodPost, tc.url, nil)
			w := httptest.NewRecorder()

			assert.Panics(t, func() {
				handler.ServeHTTP(w, req)
			}, "Expected panic for invalid input")
		})
	}
}
