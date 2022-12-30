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
	require.Error(t, Error[string]{Err: MinCharsNotReached(10, shortString)}, result)
}

func TestStringFieldMaxAndMinCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	field := String(MinChars(10), MaxChars(15))
	result := field.Validate(longString)
	require.Error(t, Error[string]{Err: MaxCharsReached(10, longString)}, result)
}
