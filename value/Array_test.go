package value_test

import (
	. "github.com/joaonrb/alsatian/result"
	. "github.com/joaonrb/alsatian/value"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayValueMaxAndMinCharShouldReturnOkStringWhenValueLengthIsBetweenMinAndMax(t *testing.T) {
	value := Array[int](Min[int](10), Max[int](15))
	result := value.Validate(mediumIntArray)
	require.Equal(t, OK[[]int]{Value: mediumIntArray}, result)
}

func TestArrayValueMaxAndMinCharShouldReturnErrorStringWhenValueHasLessThanMinChars(t *testing.T) {
	value := Array[int](Min[int](10), Max[int](15))
	result := value.Validate(shortIntArray)
	require.IsType(t, Error[[]int]{}, result)
	require.IsType(t, MinNotReachedError{}, result.(Error[[]int]).Err)
	require.Equal(t, uint64(10), result.(Error[[]int]).Err.(MinNotReachedError).Min())
	require.Equal(t, shortIntArray, result.(Error[[]int]).Err.(MinNotReachedError).Value())
}

func TestArrayValueMaxAndMinCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	value := Array[int](Min[int](10), Max[int](15))
	result := value.Validate(longIntArray)
	require.IsType(t, Error[[]int]{}, result)
	require.IsType(t, MaxReachedError{}, result.(Error[[]int]).Err)
	require.Equal(t, uint64(15), result.(Error[[]int]).Err.(MaxReachedError).Max())
	require.Equal(t, longIntArray, result.(Error[[]int]).Err.(MaxReachedError).Value())
}
