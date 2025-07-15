package common

type Identifiable[T comparable] interface {
	ID() T
}
