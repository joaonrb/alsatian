package value

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
	"reflect"
)

type Value[T any] struct {
	option Option[T]
}

func (value *Value[T]) Validate(v T) Result[T] {
	return value.option.validate(OK[T]{Value: v})
}

func (value *Value[T]) String() string {
	var target T
	return fmt.Sprintf("Value[%s]", reflect.TypeOf(target).Name())
}
