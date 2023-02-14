package result

type Result[T any] interface {
	Do(func(result T) Result[T]) Result[T]
	IfOK(func(result T)) Result[T]
	IfError(func(err error)) Result[T]
	result()
}
