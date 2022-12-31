package field

func MaybeString(options ...Option[string]) *Field[*string] {
	option := aggregate(options...)
	return &Field[*string]{
		option: &nullable[*string, string]{
			Option: option,
		},
	}
}
