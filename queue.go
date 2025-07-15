package common

type Queue[T Comparable] interface {
	Push(element T) bool
	Pop() (T, bool)
	Peek() (T, bool)
	Remove(element T) bool
	Len() int
}
