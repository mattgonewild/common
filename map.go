package common

type Map[T any] interface {
	Set(key Key, value T) error
	Get(key Key) (T, error)
	Delete(key Key) error
}
