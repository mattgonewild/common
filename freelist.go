package common

type Freelist[T any] interface {
	Get() (*T, bool)
	Release(id uint) bool
	Len() int
	Cap() int
}
