package value

func Array[T any](options ...Option[[]T]) *Value[[]T] {
	return &Value[[]T]{
		option: aggregate(options...),
	}
}
