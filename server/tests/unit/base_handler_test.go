package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"contest-influence/server/internal/handlers"
	"contest-influence/server/internal/handlers/handler_types"

	"github.com/stretchr/testify/assert"
)

type HandlerMock struct {
	PanicStruct          any
	ServeHTTPCalledCount int64
}

func (h *HandlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ServeHTTPCalledCount += 1
	if h.PanicStruct != nil {
		panic(h.PanicStruct)
	}
}

func ExpectedServeHTTPCalledNTimes(t *testing.T, expected int64, handler *HandlerMock) {
	assert.Equalf(t, expected, handler.ServeHTTPCalledCount, "Wrong number of calling ServeHTTP")
}

func TestBaseHandler_Method(t *testing.T) {
	handlerMock := HandlerMock{
		PanicStruct: nil,
	}

	handler := handlers.BaseHandler{
		Handler: &handlerMock,
		Method:  http.MethodPost,
	}

	requests := []*http.Request{
		httptest.NewRequest(http.MethodPost, "/api/v1/path?param=val", nil),
		httptest.NewRequest(http.MethodGet, "/api/v1/path?param=val", nil),
	}

	for _, r := range requests {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		ExpectedServeHTTPCalledNTimes(t, 1, &handlerMock)

		if handler.Method == r.Method {
			ExpectStatusCodesEqual(t, http.StatusOK, w.Result().StatusCode)
		} else {
			ExpectStatusCodesEqual(t, http.StatusMethodNotAllowed, w.Result().StatusCode)
		}
	}
}

type TestingError struct {
	Message string
}

func (e TestingError) Error() string {
	return e.Message
}

func TestBaseHandler_Panic(t *testing.T) {
	handlerMocks := []HandlerMock{
		{
			PanicStruct: nil,
		},
		{
			PanicStruct: handler_types.HandlerPanic{
				StatusCode: http.StatusBadRequest,
				Message:    "bad request",
			},
		},
		{
			PanicStruct: TestingError{
				Message: "bad request",
			},
		},
	}

	for _, handlerMock := range handlerMocks {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/path?param=val", nil)

		handler := handlers.BaseHandler{
			Handler: &handlerMock,
			Method:  http.MethodGet,
		}

		handler.ServeHTTP(w, r)

		ExpectedServeHTTPCalledNTimes(t, 1, &handlerMock)

		switch err := handlerMock.PanicStruct.(type) {
		case nil:
			ExpectStatusCodesEqual(t, http.StatusOK, w.Result().StatusCode)
		case handler_types.HandlerPanic:
			ExpectStatusCodesEqual(t, err.StatusCode, w.Result().StatusCode)
			assert.Contains(t, w.Body.String(), err.Message)
		default:
			ExpectStatusCodesEqual(t, http.StatusInternalServerError, w.Result().StatusCode)
			assert.Contains(t, w.Body.String(), err.(error).Error())
		}
	}
}
