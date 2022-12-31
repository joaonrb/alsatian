package field

import . "github.com/joaonrb/alsatian/result"

type pointer[T any] interface {
	*T
}

type nullable[P pointer[T], T any] struct {
	Option[T]
}

func (option *nullable[P, T]) validate(result Result[P]) Result[P] {
	return Pipe(result, func(p P) Result[P] {
		if p == nil {
			return OK[P]{Value: p}
		}
		return Pipe(option.Option.validate(OK[T]{Value: *p}), func(t T) Result[P] {
			return OK[P]{Value: &t}
		})
	})
}
