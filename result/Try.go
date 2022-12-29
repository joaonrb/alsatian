package result

import (
	goerrors "github.com/go-errors/errors"
)

func Try[V any, R any](fn func(value V) Result[R]) func(V) Result[R] {
	return func(value V) (result Result[R]) {
		defer func() {
			exception := recover()
			if exception != nil {
				result = Error[R]{
					Error: goerrors.New(exception),
				}
			}
		}()
		return fn(value)
	}
}
