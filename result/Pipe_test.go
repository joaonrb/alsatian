package result_test

import (
	"github.com/go-errors/errors"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPipeOkWithCorrectFunctionShouldReturnCorrectResult(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "1"}
	result := Pipe(ok, convert[int])
	require.Equal(t, OK[int]{Value: 1}, result)
}

func TestPipeOkWithInvalidFunctionShouldReturnErrorResult(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "J"}
	result := Pipe(ok, convert[int])
	require.IsType(t, Error[int]{}, result)
}

func TestPipeErrorWithCorrectFunctionShouldReturnErrorResult(t *testing.T) {
	var err Result[string] = Error[string]{Error: errors.New("something dark side")}
	result := Pipe(err, convert[int])
	require.Equal(t, Error[int]{Error: err.(Error[string]).Error}, result)
}
