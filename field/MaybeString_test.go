package field_test

import (
	. "github.com/joaonrb/alsatian/field"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaybeStringFieldMaxAndMinCharShouldReturnOkStringWhenValueIsNil(t *testing.T) {
	field := MaybeString(MinChars(10), MaxChars(15))
	result := field.Validate(nil)
	require.Equal(t, OK[*string]{Value: nil}, result)
}

func TestMaybeStringFieldMaxAndMinCharShouldReturnOkStringWhenValueLengthIsBetweenMinAndMax(t *testing.T) {
	str := mediumString
	field := MaybeString(MinChars(10), MaxChars(15))
	result := field.Validate(&str)
	require.Equal(t, OK[*string]{Value: &str}, result)
}

func TestMaybeStringFieldMaxAndMinCharShouldReturnErrorStringWhenValueHasLessThanMinChars(t *testing.T) {
	str := shortString
	field := MaybeString(MinChars(10), MaxChars(15))
	result := field.Validate(&str)
	require.IsType(t, Error[*string]{}, result)
	require.IsType(t, MinCharsNotReachedError{}, result.(Error[*string]).Err)
	require.Equal(t, uint64(10), result.(Error[*string]).Err.(MinCharsNotReachedError).Min())
	require.Equal(t, shortString, result.(Error[*string]).Err.(MinCharsNotReachedError).Value())
}

func TestMaybeStringFieldMaxAndMinCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	str := longString
	field := MaybeString(MinChars(10), MaxChars(15))
	result := field.Validate(&str)
	require.IsType(t, Error[*string]{}, result)
	require.IsType(t, MaxCharsReachedError{}, result.(Error[*string]).Err)
	require.Equal(t, uint64(15), result.(Error[*string]).Err.(MaxCharsReachedError).Max())
	require.Equal(t, longString, result.(Error[*string]).Err.(MaxCharsReachedError).Value())
}
