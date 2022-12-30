package field

func String(options ...Option[string]) *Field[string] {
	return &Field[string]{
		option: aggregate(options...),
	}
}
