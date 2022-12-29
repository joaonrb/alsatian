package result_test

import (
	. "github.com/joaonrb/alsatian/result"
	"strconv"
)

func convertInt(value string) Result[int] {
	n, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return Error[int]{Error: err}
	}
	return OK[int]{Value: int(n)}
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func divideBy[N Number](n int64) func(N) Result[int64] {
	return func(value N) Result[int64] {
		return OK[int64]{Value: int64(value) / n}
	}
}
