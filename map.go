package common

type Map[K, V any] interface {
	Set(key K, value V) error
	Get(key K) (V, error)
	Has(key K) bool
	Delete(key K) error
	ForEach(yield func(K, V) bool)
}
