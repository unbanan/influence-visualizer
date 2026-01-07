package unit

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"contest-influence/server/internal/handlers"

	"github.com/stretchr/testify/assert"
)

type RegisterImplMock struct {
	RegisterCalledCount int64
	LastID              int64
	LastName            string
	ShouldPanic         bool
	PanicMessage        string
}

func (m *RegisterImplMock) Register(id int64, name string) {
	m.RegisterCalledCount++
	m.LastID = id
	m.LastName = name
	if m.ShouldPanic {
		panic(m.PanicMessage)
	}
}

func TestRegisterHandler_ValidRequest(t *testing.T) {
	mock := &RegisterImplMock{}
	regex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)

	handler := &handlers.RegisterHandler{
		Impl:  mock,
		Regex: regex,
	}

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
	mock := &RegisterImplMock{}
	regex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)

	handler := &handlers.RegisterHandler{
		Impl:  mock,
		Regex: regex,
	}

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
			mock.RegisterCalledCount = 0
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/api/v1/register"+tc.queryStr, nil)

			handler.ServeHTTP(w, r)

			assert.Equal(t, int64(0), mock.RegisterCalledCount)
			ExpectStatusCodesEqual(t, http.StatusBadRequest, w.Result().StatusCode)
			assert.Contains(t, w.Body.String(), "Invalid name")
		})
	}
}

func TestRegisterHandler_ImplPanic(t *testing.T) {
	mock := &RegisterImplMock{
		ShouldPanic:  true,
		PanicMessage: "database error",
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)

	handler := &handlers.RegisterHandler{
		Impl:  mock,
		Regex: regex,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/v1/register?id=123&name=testuser", nil)

	assert.Panics(t, func() {
		handler.ServeHTTP(w, r)
	})
	assert.Equal(t, int64(1), mock.RegisterCalledCount)
}
