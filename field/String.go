package field

func String(options ...Option[string]) *Field[string] {
	field := &Field[string]{}
	field.init(options...)
	return field
}
