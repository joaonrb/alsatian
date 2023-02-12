package value_test

import (
	. "github.com/joaonrb/alsatian/result"
	. "github.com/joaonrb/alsatian/value"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringValueMaxAndMinCharShouldReturnOkStringWhenValueLengthIsBetweenMinAndMax(t *testing.T) {
	value := String(MinChars(10), MaxChar(15))
	result := value.Validate(mediumString)
	require.Equal(t, OK[string]{Value: mediumString}, result)
}

func TestStringValueMaxAndMinCharShouldReturnErrorStringWhenValueHasLessThanMinChars(t *testing.T) {
	value := String(MinChars(10), MaxChar(15))
	result := value.Validate(shortString)
	require.IsType(t, Error[string]{}, result)
	require.IsType(t, MinNotReachedError{}, result.(Error[string]).Err)
	require.Equal(t, uint64(10), result.(Error[string]).Err.(MinNotReachedError).Min())
	require.Equal(t, shortString, result.(Error[string]).Err.(MinNotReachedError).Value())
}

func TestStringValueMaxAndMinCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	value := String(MinChars(10), MaxChar(15))
	result := value.Validate(longString)
	require.IsType(t, Error[string]{}, result)
	require.IsType(t, MaxReachedError{}, result.(Error[string]).Err)
	require.Equal(t, uint64(15), result.(Error[string]).Err.(MaxReachedError).Max())
	require.Equal(t, longString, result.(Error[string]).Err.(MaxReachedError).Value())
}
