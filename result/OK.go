package result

import (
	"fmt"
	"reflect"
)

type OK[T any] struct {
	Value T
}

func (ok OK[T]) withValue(fn func(value T)) {
	fn(ok.Value)
}

func (ok OK[T]) withError(fn func(error)) {}

func (ok OK[T]) String() string {
	return fmt.Sprintf("result %s OK: %v", reflect.TypeOf(ok.Value).Name(), ok.Value)
}
