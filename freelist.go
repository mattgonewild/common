package common

type Freelist[T Identifiable[uint]] interface {
	Get() (*T, bool)
	Release(element *T) bool
	Len() int
	Cap() int
}
