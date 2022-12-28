package result

type Result[T any] interface {
	withValue(func(result T))
	withError(func(err error))
}
