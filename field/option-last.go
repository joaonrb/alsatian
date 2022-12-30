package field

import . "github.com/joaonrb/alsatian/result"

type last[T any] struct{}

func (*last[T]) validate(result Result[T]) Result[T] {
	return result
}

func (option *last[T]) String() string {
	return "LastOption"
}
