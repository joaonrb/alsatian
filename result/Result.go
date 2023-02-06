package result

type Result[T any] interface {
	IfOK(func(result T))
	IfError(func(err error))
	result()
}
