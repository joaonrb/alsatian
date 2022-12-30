package field

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
)

func MinChars(Min uint64) Option[string] {
	return &node[string]{
		option: &minChars{min: Min},
	}
}

type minChars struct {
	min uint64
}

func (option *minChars) validate(result Result[string]) Result[string] {
	return Pipe(result, func(value string) Result[string] {
		if len(value) < int(option.min) {
			return Error[string]{Err: MinCharsNotReached(option.min, value)}
		}
		return result
	})
}

func (option *minChars) String() string {
	return fmt.Sprintf("MinChars(%d)", option.min)
}
