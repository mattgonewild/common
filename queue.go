package common

type Queue[T any] interface {
	Push(element T)
	Pop() (T, bool)
	Peek() (T, bool)
	Len() int
}
