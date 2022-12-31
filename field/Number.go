package field

import "golang.org/x/exp/constraints"

func Number[N constraints.Signed | constraints.Float](options ...Option[N]) *Field[N] {
	return &Field[N]{
		option: aggregate(options...),
	}
}
