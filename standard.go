package common

type Standard[T any, I comparable] interface {
	Comparable[T]
	Identifiable[I]
}
