package field

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
	"reflect"
)

type Field[T any] struct {
	option Option[T]
}

func (field *Field[T]) Validate(value T) Result[T] {
	return field.option.validate(OK[T]{Value: value})
}

func (field *Field[T]) String() string {
	var target T
	return fmt.Sprintf("Field%s", reflect.TypeOf(target).Name())
}

func aggregate[T any](options ...Option[T]) Option[T] {
	options = append(options, &last[T]{})
	head := options[0]
	for _, option := range options[1:] {
		head.integrate(option)
	}
	return head
}
