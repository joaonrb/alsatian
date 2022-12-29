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

func (n *node[T]) integrate(o Option[T]) {
	if n.next == nil {
		n.next = o
	} else {
		n.next.integrate(o)
	}
}

func (n *node[T]) String() string {
	return fmt.Sprintf("%s", n.option)
}
