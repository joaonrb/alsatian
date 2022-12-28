package result_test

import (
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func convertInt(value string) Result[int] {
	n, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return Error[int]{Error: err}
	}
	return OK[int]{Value: int(n)}
}

func TestPipeOkWithCorrectFunctionShouldReturnCorrectResult(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "1"}
	result := Pipe(ok, convertInt)
	assert.Equal(t, OK[int]{Value: 1}, result)
}

func TestPipeOkWithInvalidFunctionShouldReturnErrorResult(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "J"}
	result := Pipe(ok, convertInt)
	assert.IsType(t, Error[int]{}, result)
}
