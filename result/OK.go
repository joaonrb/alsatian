package result

import (
	"fmt"
	"reflect"
)

type OK[T any] struct {
	Value T
}

func (ok OK[T]) IfOK(fn func(value T) Result[T]) Result[T] {
	return fn(ok.Value)
}

func (ok OK[T]) IfError(fn func(error) Result[T]) Result[T] { return ok }

func (ok OK[T]) String() string {
	return fmt.Sprintf("result %s OK: %v", reflect.TypeOf(ok.Value).Name(), ok.Value)
}

func (ok OK[T]) result() {}
