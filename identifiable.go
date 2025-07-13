package common

type Identifiable[T comparable] interface {
	ID() T
	Equal(identifiable Identifiable[T]) bool
}
