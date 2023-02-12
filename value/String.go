package value

func String(options ...Option[string]) *Value[string] {
	return &Value[string]{
		option: aggregate(options...),
	}
}
