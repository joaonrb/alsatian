package value

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
)

func MinChars(Min uint64) Option[string] {
	return &bytesToStringOption{Option: &min[[]byte, byte]{len: Min}}
}

func Min[T any](Min uint64) Option[[]T] {
	return &min[[]T, T]{len: Min}
}

type min[A []T, T any] struct {
	len uint64
}

func (option *min[A, T]) validate(result Result[A]) Result[A] {
	return IfOK(result, func(value A) Result[A] {
		if len(value) < int(option.len) {
			return Error[A]{Err: MinNotReached(option.len, value)}
		}
		return result
	})
}

func (option *min[A, T]) String() string {
	return fmt.Sprintf("Min(%d)", option.len)
}
