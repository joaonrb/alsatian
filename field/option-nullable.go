package field

import . "github.com/joaonrb/alsatian/result"

type nullable[T any] struct {
	Option[T]
}

func (option *nullable[T]) validate(result Result[*T]) Result[*T] {
	return Pipe(result, func(p *T) Result[*T] {
		if p == nil {
			return OK[*T]{Value: p}
		}
		return Pipe(option.Option.validate(OK[T]{Value: *p}), func(t T) Result[*T] {
			return OK[*T]{Value: &t}
		})
	})
}
