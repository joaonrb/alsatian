package field

import (
	"fmt"
	"github.com/go-errors/errors"
	"golang.org/x/exp/constraints"
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

func (err MaxCharsReachedError) Max() uint64 {
	return err.max
}

func (err MaxCharsReachedError) Value() string {
	return err.value
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

func (err MinCharsNotReachedError) Min() uint64 {
	return err.min
}

func (err MinCharsNotReachedError) Value() string {
	return err.value
}

func ValueHigherThan[N constraints.Signed | constraints.Float](value N, limit N) ValueHigherThanError[N] {
	return ValueHigherThanError[N]{
		inner: errors.Wrap(fmt.Errorf("value %v is bigger than %v", value, limit), 1),
		value: value,
		limit: limit,
	}
}

type ValueHigherThanError[N constraints.Signed | constraints.Float] struct {
	inner
	value N
	limit N
}

func (err ValueHigherThanError[N]) Value() N {
	return err.value
}

func (err ValueHigherThanError[N]) Limit() N {
	return err.limit
}
