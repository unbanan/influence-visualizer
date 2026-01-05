package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExpectStatusCodesEqual(t *testing.T, expected, actual int) {
	assert.Equalf(t, expected, actual, "Wrong status code")
}
