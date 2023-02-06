package result_test

import (
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTryWithCorrectFunctionShouldReturnOk(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "1"}
	result := IfOK(ok, Try(convert[int]))
	require.Equal(t, OK[int]{Value: 1}, result)
}

func TestTryWithInvalidFunctionShouldReturnError(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "1"}
	result := IfOK(
		IfOK(ok, Try(convert[int])),
		Try(divideBy[int](0)),
	)
	require.IsType(t, Error[int]{}, result)
}
