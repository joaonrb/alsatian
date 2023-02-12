package value

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
	"golang.org/x/exp/constraints"
)

func LowerThan[N constraints.Signed | constraints.Float](Limit N) Option[N] {
	return &lowerThan[N]{limit: Limit}
}

type lowerThan[N constraints.Signed | constraints.Float] struct {
	limit N
}

func (option *lowerThan[N]) validate(result Result[N]) Result[N] {
	return IfOK(result, func(value N) Result[N] {
		if value > option.limit {
			return Error[N]{Err: HigherThan(value, option.limit)}
		}
		return result
	})
}

func (option *lowerThan[N]) String() string {
	return fmt.Sprintf("LowerThan(%v)", option.limit)
}
