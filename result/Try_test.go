package result_test

import (
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryWithCorrectFunctionShouldReturnOk(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "1"}
	result := Pipe(ok, Try(convertInt))
	assert.Equal(t, OK[int]{Value: 1}, result)
}

func TestTryWithInvalidFunctionShouldReturnError(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "1"}
	result := Pipe(Pipe(ok, Try(convertInt)), Try(divideBy[int](0)))
	assert.IsType(t, Error[int64]{}, result)
}
