package params

import (
	"fmt"
	"net/http"
	"strconv"

	"contest-influence/server/internal/handlers/handler_types"
)

func getField(r *http.Request, field string) string {
	if !r.URL.Query().Has(field) {
		panic(handler_types.HandlerPanic{
			Message:    "",
			StatusCode: http.StatusBadRequest,
		})
	}

	return r.URL.Query().Get(field)
}

func GetString(r *http.Request, field string) string {
	return getField(r, field)
}

func GetInt(r *http.Request, field string) int64 {
	val, err := strconv.ParseInt(getField(r, field), 10, 64)

	if err != nil {
		panic(handler_types.HandlerPanic{
			Message:    fmt.Sprintf("field %s has wrong format (should be an int)", field),
			StatusCode: http.StatusBadRequest,
		})
	}

	return val
}
