package value

import . "github.com/joaonrb/alsatian/result"

type nullable[T any] struct {
	Option[T]
}

func (option *nullable[T]) validate(result Result[*T]) Result[*T] {
	return IfOK(result, func(p *T) Result[*T] {
		if p == nil {
			return OK[*T]{Value: p}
		}
		return IfOK(option.Option.validate(OK[T]{Value: *p}), func(t T) Result[*T] {
			return OK[*T]{Value: &t}
		})
	})
}
