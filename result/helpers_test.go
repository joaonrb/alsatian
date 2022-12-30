package result_test

import (
	. "github.com/joaonrb/alsatian/result"
	"golang.org/x/exp/constraints"
	"strconv"
)

func convert[T constraints.Integer](value string) Result[T] {
	n, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return Error[T]{Err: err}
	}
	return OK[T]{Value: T(n)}
}

func divideBy[N constraints.Integer](n N) func(N) Result[N] {
	return func(value N) Result[N] {
		return OK[N]{Value: value / n}
	}
}
