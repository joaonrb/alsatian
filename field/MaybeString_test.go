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
	require.Error(t, Error[string]{Err: MinCharsNotReached(10, str)}, result)
}

func TestMaybeStringFieldMaxAndMinCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	str := longString
	field := MaybeString(MinChars(10), MaxChars(15))
	result := field.Validate(&str)
	require.Error(t, Error[string]{Err: MaxCharsReached(10, str)}, result)
}
