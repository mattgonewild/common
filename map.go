package common

type Map[K, V any] interface {
	Set(key K, value V) error
	Get(key K) (V, error)
	Delete(key K) error
	ForEach(yield func(K, V) bool)
}
