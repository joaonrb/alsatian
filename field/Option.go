package field

import . "github.com/joaonrb/alsatian/result"

type Option[T any] interface {
	validate(result Result[T]) Result[T]
}
