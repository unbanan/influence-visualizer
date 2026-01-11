package unit

import (
	"errors"
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
			PanicStruct: errors.New("bad request"),
		},
	}

	for _, handlerMock := range handlerMocks {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/v1/path?param=val", nil)

		handler := handlers.WrapHandler(&handlerMock)

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
