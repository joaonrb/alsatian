package result

func IfOK[V any, R any](value Result[V], fn func(V) Result[R]) (result Result[R]) {
	value.IfError(func(err error) Result[V] {
		result = Error[R]{Err: err}
		return value
	})
	value.IfOK(func(v V) Result[V] {
		result = fn(v)
		return value
	})
	return
}
