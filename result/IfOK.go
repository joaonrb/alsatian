package result

func IfOK[V any, R any](value Result[V], fn func(V) Result[R]) (result Result[R]) {
	value.IfError(func(err error) {
		result = Error[R]{Err: err}
	})
	value.IfOK(func(v V) {
		result = fn(v)
	})
	return
}
