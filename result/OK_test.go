package result_test

import (
	"github.com/go-errors/errors"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOKIfOKWithCorrectFunctionShouldReturnCorrectResult(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "1"}
	var result Result[int]
	ok.IfOK(func(value string) {
		result = convert[int](value)
	})
	require.Equal(t, OK[int]{Value: 1}, result)
}

func TestOKIfOKWithInvalidFunctionShouldReturnErrorResult(t *testing.T) {
	var ok Result[string] = OK[string]{Value: "J"}
	var result Result[int]
	ok.IfOK(func(value string) {
		result = convert[int](value)
	})
	require.IsType(t, Error[int]{}, result)
}

func TestErrIfOKWithCorrectFunctionShouldReturnErrorResult(t *testing.T) {
	var err Result[string] = Error[string]{Err: errors.New("something dark side")}
	var result Result[int]
	err.IfOK(func(value string) {
		result = convert[int](value)
	})
	require.Equal(t, nil, result)
}
