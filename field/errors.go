package field

import (
	"fmt"
	"github.com/go-errors/errors"
)

type inner interface {
	Error() string
	ErrorStack() string
	StackFrames() []errors.StackFrame
}

func MaxCharsReached(max uint64, value string) MaxCharsReachedError {
	return MaxCharsReachedError{
		inner: errors.Wrap(fmt.Errorf("max of %d chars reached for value '%s'", max, value), 1),
		value: value,
		max:   max,
	}
}

type MaxCharsReachedError struct {
	inner
	value string
	max   uint64
}

func MinCharsNotReached(min uint64, value string) MinCharsNotReachedError {
	return MinCharsNotReachedError{
		inner: errors.Wrap(fmt.Errorf("min of %d chars not reached for value '%s'", min, value), 1),
		value: value,
		min:   min,
	}
}

type MinCharsNotReachedError struct {
	inner
	value string
	min   uint64
}
