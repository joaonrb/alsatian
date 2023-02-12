package value_test

import (
	. "github.com/joaonrb/alsatian/result"
	. "github.com/joaonrb/alsatian/value"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringValueMinCharShouldReturnOkStringWhenValueHasAtLeastMinChars(t *testing.T) {
	value := String(MinChars(10))
	result := value.Validate(longString)
	require.Equal(t, OK[string]{Value: longString}, result)
}

func TestStringValueMinCharShouldReturnErrorStringWhenValueHasLessThanMinChars(t *testing.T) {
	value := String(MinChars(10))
	result := value.Validate(shortString)
	require.IsType(t, Error[string]{}, result)
	require.IsType(t, MinNotReachedError{}, result.(Error[string]).Err)
	require.Equal(t, uint64(10), result.(Error[string]).Err.(MinNotReachedError).Min())
	require.Equal(t, shortString, result.(Error[string]).Err.(MinNotReachedError).Value())

}
