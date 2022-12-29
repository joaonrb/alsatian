package field_test

import (
	. "github.com/joaonrb/alsatian/field"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringFieldMinCharShouldReturnOkStringWhenValueHasAtLeastMinChars(t *testing.T) {
	field := String(MinChars(10))
	result := field.Validate("Joao Nuno Ramalho Baptista")
	require.IsType(t, OK[string]{}, result)
	value := result.(OK[string]).Value
	require.Equal(t, "Joao Nuno Ramalho Baptista", value)
}

func TestStringFieldMinCharShouldReturnErrorStringWhenValueHasLessThanMinChars(t *testing.T) {
	field := String(MinChars(10))
	result := field.Validate("Joao")
	require.IsType(t, Error[string]{}, result)
}
