package common

type Freelist[T Comparable[T]] interface {
	Get() (T, bool)
	Release(element T) bool
	Len() int
	Cap() int
}
