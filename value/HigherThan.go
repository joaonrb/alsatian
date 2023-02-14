package value

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
	"golang.org/x/exp/constraints"
)

func HigherThan[N constraints.Integer | constraints.Float](Limit N) Option[N] {
	return &higherThan[N]{limit: Limit}
}

type higherThan[N constraints.Integer | constraints.Float] struct {
	limit N
}

func (option *higherThan[N]) validate(result Result[N]) Result[N] {
	return IfOK(result, func(value N) Result[N] {
		if value < option.limit {
			return Error[N]{Err: ValueLowerThan(value, option.limit)}
		}
		return result
	})
}

func (option *higherThan[N]) String() string {
	return fmt.Sprintf("HigherThan(%v)", option.limit)
}
