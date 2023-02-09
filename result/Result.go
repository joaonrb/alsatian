package result

type Result[T any] interface {
	IfOK(func(result T) Result[T]) Result[T]
	IfError(func(err error) Result[T]) Result[T]
	result()
}
