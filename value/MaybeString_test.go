package value_test

import (
	. "github.com/joaonrb/alsatian/result"
	. "github.com/joaonrb/alsatian/value"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaybeStringValueMaxAndMinCharShouldReturnOkStringWhenValueIsNil(t *testing.T) {
	value := MaybeString(MinChars(10), MaxChar(15))
	result := value.Validate(nil)
	require.Equal(t, OK[*string]{Value: nil}, result)
}

func TestMaybeStringValueMaxAndMinCharShouldReturnOkStringWhenValueLengthIsBetweenMinAndMax(t *testing.T) {
	str := mediumString
	value := MaybeString(MinChars(10), MaxChar(15))
	result := value.Validate(&str)
	require.Equal(t, OK[*string]{Value: &str}, result)
}

func TestMaybeStringValueMaxAndMinCharShouldReturnErrorStringWhenValueHasLessThanMinChars(t *testing.T) {
	str := shortString
	value := MaybeString(MinChars(10), MaxChar(15))
	result := value.Validate(&str)
	require.IsType(t, Error[*string]{}, result)
	require.IsType(t, MinNotReachedError{}, result.(Error[*string]).Err)
	require.Equal(t, uint64(10), result.(Error[*string]).Err.(MinNotReachedError).Min())
	require.Equal(t, shortString, result.(Error[*string]).Err.(MinNotReachedError).Value())
}

func TestMaybeStringValueMaxAndMinCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	str := longString
	value := MaybeString(MinChars(10), MaxChar(15))
	result := value.Validate(&str)
	require.IsType(t, Error[*string]{}, result)
	require.IsType(t, MaxReachedError{}, result.(Error[*string]).Err)
	require.Equal(t, uint64(15), result.(Error[*string]).Err.(MaxReachedError).Max())
	require.Equal(t, longString, result.(Error[*string]).Err.(MaxReachedError).Value())
}
