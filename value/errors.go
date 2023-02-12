package value

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

type base struct {
	value any
}

func (err *base) Value() any {
	return err.value
}

func (err *base) setValue(value any) {
	err.value = value
}

func MaxReached(max uint64, value any) MaxReachedError {
	return MaxReachedError{
		inner: errors.Wrap(fmt.Errorf("len of %d reached for value '%s'", max, value), 1),
		base:  &base{value: value},
		max:   max,
	}
}

type MaxReachedError struct {
	inner
	*base
	max uint64
}

func (err MaxReachedError) Max() uint64 {
	return err.max
}

func MinNotReached(min uint64, value any) MinNotReachedError {
	return MinNotReachedError{
		inner: errors.Wrap(fmt.Errorf("len of %d chars not reached for value '%s'", min, value), 1),
		base:  &base{value: value},
		min:   min,
	}
}

type MinNotReachedError struct {
	inner
	*base
	min uint64
}

func (err MinNotReachedError) Min() uint64 {
	return err.min
}

func HigherThan[N constraints.Signed | constraints.Float](value N, limit N) HigherThanError[N] {
	return HigherThanError[N]{
		inner: errors.Wrap(fmt.Errorf("value %v is bigger than %v", value, limit), 1),
		base:  &base{value: value},
		limit: limit,
	}
}

type HigherThanError[N constraints.Signed | constraints.Float] struct {
	inner
	*base
	limit N
}

func (err HigherThanError[N]) Limit() N {
	return err.limit
}
