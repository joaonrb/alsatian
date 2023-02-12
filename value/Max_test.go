package value_test

import (
	. "github.com/joaonrb/alsatian/result"
	. "github.com/joaonrb/alsatian/value"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringValueMaxCharShouldReturnOkStringWhenValueHasLessThanMaxChars(t *testing.T) {
	value := String(MaxChar(10))
	result := value.Validate(shortString)
	require.Equal(t, OK[string]{Value: shortString}, result)
}

func TestStringValueMaxCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	value := String(MaxChar(10))
	result := value.Validate(longString)
	require.IsType(t, Error[string]{}, result)
	require.IsType(t, MaxReachedError{}, result.(Error[string]).Err)
	require.Equal(t, uint64(10), result.(Error[string]).Err.(MaxReachedError).Max())
	require.Equal(t, longString, result.(Error[string]).Err.(MaxReachedError).Value())
}
