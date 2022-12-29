package result

func Pipe[V any, R any](value Result[V], fn func(V) Result[R]) (result Result[R]) {
	value.withError(func(err error) {
		result = Error[R]{Error: err}
	})
	value.withValue(func(v V) {
		result = fn(v)
	})
	return
}
