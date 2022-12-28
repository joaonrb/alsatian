package result

func Pipe[T1 any, T2 any](value Result[T1], fn func(T1) Result[T2]) (result Result[T2]) {
	value.withError(func(err error) {
		result = Error[T2]{Error: err}
	})
	value.withValue(func(v T1) {
		result = fn(v)
	})
	return
}
