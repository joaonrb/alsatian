package field_test

import (
	. "github.com/joaonrb/alsatian/field"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringFieldMaxAndMinCharShouldReturnOkStringWhenValueLengthIsBetweenMinAndMax(t *testing.T) {
	field := String(MinChars(10), MaxChars(15))
	result := field.Validate(mediumString)
	require.Equal(t, OK[string]{Value: mediumString}, result)
}

func TestStringFieldMaxAndMinCharShouldReturnErrorStringWhenValueHasLessThanMinChars(t *testing.T) {
	field := String(MinChars(10), MaxChars(15))
	result := field.Validate(shortString)
	require.IsType(t, Error[string]{}, result)
	require.IsType(t, MinCharsNotReachedError{}, result.(Error[string]).Err)
	require.Equal(t, uint64(10), result.(Error[string]).Err.(MinCharsNotReachedError).Min())
	require.Equal(t, shortString, result.(Error[string]).Err.(MinCharsNotReachedError).Value())
}

func TestStringFieldMaxAndMinCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	field := String(MinChars(10), MaxChars(15))
	result := field.Validate(longString)
	require.IsType(t, Error[string]{}, result)
	require.IsType(t, MaxCharsReachedError{}, result.(Error[string]).Err)
	require.Equal(t, uint64(15), result.(Error[string]).Err.(MaxCharsReachedError).Max())
	require.Equal(t, longString, result.(Error[string]).Err.(MaxCharsReachedError).Value())
}
