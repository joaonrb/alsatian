package result

import (
	"fmt"
	"reflect"
)

type Error[T any] struct {
	Error error
}

func (err Error[T]) withValue(func(T)) {}

func (err Error[T]) withError(fn func(error)) {
	fn(err.Error)
}

func (err Error[T]) String() string {
	var target T
	return fmt.Sprintf("result %s error: %s", reflect.TypeOf(target).Name(), err.Error)
}
