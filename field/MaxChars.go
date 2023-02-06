package field

import (
	"fmt"
	. "github.com/joaonrb/alsatian/result"
)

func MaxChars(max uint64) Option[string] {
	return &maxChars{max: max}
}

type maxChars struct {
	max uint64
}

func (option *maxChars) validate(result Result[string]) Result[string] {
	return IfOK(result, func(value string) Result[string] {
		if len(value) > int(option.max) {
			return Error[string]{Err: MaxCharsReached(option.max, value)}
		}
		return result
	})
}

func (option *maxChars) String() string {
	return fmt.Sprintf("MaxChars(%d)", option.max)
}
