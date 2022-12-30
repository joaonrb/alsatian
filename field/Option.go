package field

type Option[T any] interface {
	validatable[T]
	integrate(Option[T])
}
