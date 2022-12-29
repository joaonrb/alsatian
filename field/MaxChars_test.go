package field_test

import (
	. "github.com/joaonrb/alsatian/field"
	. "github.com/joaonrb/alsatian/result"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringFieldMaxCharShouldReturnOkStringWhenValueHasLessThanMaxChars(t *testing.T) {
	field := String(MaxChars(10))
	result := field.Validate("Joao")
	require.IsType(t, OK[string]{}, result)
	value := result.(OK[string]).Value
	require.Equal(t, "Joao", value)
}

func TestStringFieldMaxCharShouldReturnErrorStringWhenValueHasMoreThanMaxChars(t *testing.T) {
	field := String(MaxChars(10))
	result := field.Validate("Joao Nuno Ramalho Baptista")
	require.IsType(t, Error[string]{}, result)
}
