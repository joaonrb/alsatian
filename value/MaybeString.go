package value

func MaybeString(options ...Option[string]) *Value[*string] {
	option := aggregate(options...)
	return &Value[*string]{
		option: &nullable[string]{
			Option: option,
		},
	}
}
