package field

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
)

type node[T any] struct {
	next   Option[T]
	option Option[T]
}

func (n *node[T]) validate(result Result[T]) Result[T] {
	return n.next.validate(n.option.validate(result))
}

func (n *node[T]) String() string {
	return fmt.Sprintf("%s", n.option)
}

func aggregate[T any](options ...Option[T]) Option[T] {
	var head Option[T] = &last[T]{}
	current := head
	for i := range options {
		current = &node[T]{next: options[len(options)-(i+1)], option: current}
	}
	return head
}
