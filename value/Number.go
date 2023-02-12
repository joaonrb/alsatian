package value

import "golang.org/x/exp/constraints"

func Number[N constraints.Signed | constraints.Float](options ...Option[N]) *Value[N] {
	return &Value[N]{
		option: aggregate(options...),
	}
}
