package value

import . "github.com/joaonrb/alsatian/result"

type Option[T any] interface {
	validate(result Result[T]) Result[T]
}
