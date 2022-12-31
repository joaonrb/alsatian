package field_test

import (
	. "github.com/joaonrb/alsatian/field"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberFieldLowerThanShouldReturnOkStringWhenValueIsLowerThanValue(t *testing.T) {
	field := Number(LowerThan(10))
	result := field.Validate(smallPositiveNumber)
	require.Equal(t, OK[int]{Value: smallPositiveNumber}, result)
}

func TestNumberFieldLowerThanShouldReturnErrorStringWhenValueIsLowerThanValue(t *testing.T) {
	field := Number(LowerThan(10))
	result := field.Validate(bigPositiveNumber)
	require.IsType(t, Error[int]{}, result)
	require.IsType(t, ValueHigherThanError[int]{}, result.(Error[int]).Err)
	require.Equal(t, 10, result.(Error[int]).Err.(ValueHigherThanError[int]).Limit())
	require.Equal(t, bigPositiveNumber, result.(Error[int]).Err.(ValueHigherThanError[int]).Value())

}
