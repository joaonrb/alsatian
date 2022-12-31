package field_test

import (
	. "github.com/joaonrb/alsatian/field"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringFieldMaxCharShouldReturnOkStringWhenValueHasLessThanMaxChars(t *testing.T) {
	field := String(MaxChars(10))
	result := field.Validate(shortString)
	require.Equal(t, OK[string]{Value: shortString}, result)
}

func TestStringFieldMaxCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	field := String(MaxChars(10))
	result := field.Validate(longString)
	require.IsType(t, Error[string]{}, result)
	require.IsType(t, MaxCharsReachedError{}, result.(Error[string]).Err)
	require.Equal(t, uint64(10), result.(Error[string]).Err.(MaxCharsReachedError).Max())
	require.Equal(t, longString, result.(Error[string]).Err.(MaxCharsReachedError).Value())
}
