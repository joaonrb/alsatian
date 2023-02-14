package result

import (
	"fmt"
	"reflect"
)

type Error[T any] struct {
	Err error
}

func (err Error[T]) Error() string {
	return err.Err.Error()
}

func (err Error[T]) Do(func(T) Result[T]) Result[T] { return err }

func (err Error[T]) IfOK(func(T)) Result[T] { return err }

func (err Error[T]) IfError(fn func(error)) Result[T] {
	fn(err.Err)
	return err
}

func (err Error[T]) String() string {
	var target T
	return fmt.Sprintf("result %s error: %s", reflect.TypeOf(target).Name(), err.Err)
}

func (err Error[T]) result() {}
