package common

type Set[T Comparable[T]] interface {
	Put(element T) error
	Get(element T) (T, error)
	Delete(element T) error
	ForEach(yield func(T) bool)
}
