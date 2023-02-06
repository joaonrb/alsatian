package result_test

import (
	"github.com/go-errors/errors"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIfOkWithCorrectFunctionShouldReturnCorrectResult(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "1"}
	result := IfOK(ok, convert[int])
	require.Equal(t, OK[int]{Value: 1}, result)
}

func TestIfOkWithInvalidFunctionShouldReturnErrorResult(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "J"}
	result := IfOK(ok, convert[int])
	require.IsType(t, Error[int]{}, result)
}

func TestIfErrorWithCorrectFunctionShouldReturnErrorResult(t *testing.T) {
	var err Result[string] = Error[string]{Err: errors.New("something dark side")}
	result := IfOK(err, convert[int])
	require.Equal(t, Error[int]{Err: err.(Error[string]).Err}, result)
}
