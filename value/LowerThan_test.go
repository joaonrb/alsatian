package value_test

import (
	. "github.com/joaonrb/alsatian/result"
	. "github.com/joaonrb/alsatian/value"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberValueLowerThanShouldReturnOkStringWhenValueIsLowerThanValue(t *testing.T) {
	value := Number(LowerThan(10))
	result := value.Validate(smallPositiveNumber)
	require.Equal(t, OK[int]{Value: smallPositiveNumber}, result)
}

func TestNumberValueLowerThanShouldReturnErrorStringWhenValueIsLowerThanValue(t *testing.T) {
	value := Number(LowerThan(10))
	result := value.Validate(bigPositiveNumber)
	require.IsType(t, Error[int]{}, result)
	require.IsType(t, HigherThanError[int]{}, result.(Error[int]).Err)
	require.Equal(t, 10, result.(Error[int]).Err.(HigherThanError[int]).Limit())
	require.Equal(t, bigPositiveNumber, result.(Error[int]).Err.(HigherThanError[int]).Value())

}
