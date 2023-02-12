package value

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
)

func MaxChar(len uint64) Option[string] {
	return &bytesToStringOption{Option: &max[[]byte, byte]{len: len}}
}

func Max[T any](len uint64) Option[[]T] {
	return &max[[]T, T]{len: len}
}

type max[A []T, T any] struct {
	len uint64
}

func (option *max[A, T]) validate(result Result[A]) Result[A] {
	return IfOK(result, func(value A) Result[A] {
		if len(value) > int(option.len) {
			return Error[A]{Err: MaxReached(option.len, value)}
		}
		return result
	})
}

func (option *max[A, T]) String() string {
	return fmt.Sprintf("Max(%d)", option.len)
}
