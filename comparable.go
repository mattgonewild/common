package common

type Comparable[T any] interface {
	Before(comparable T) bool
	Equal(comparable T) bool
	After(comparable T) bool
	Compare(comparable T) int
}
