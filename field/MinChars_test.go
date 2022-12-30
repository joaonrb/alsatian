package field_test

import (
	. "github.com/joaonrb/alsatian/field"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringFieldMinCharShouldReturnOkStringWhenValueHasAtLeastMinChars(t *testing.T) {
	field := String(MinChars(10))
	result := field.Validate(longString)
	require.Equal(t, OK[string]{Value: longString}, result)
}

func TestStringFieldMinCharShouldReturnErrorStringWhenValueHasLessThanMinChars(t *testing.T) {
	field := String(MinChars(10))
	result := field.Validate(shortString)
	require.Error(t, Error[string]{Err: MinCharsNotReached(10, shortString)}, result)

}
