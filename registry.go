package common

type Registry[K comparable, V any] interface {
	Claim(key K) (V, bool)
	Register(key K, value V) error
	Release(key K) error
	Get(key K) (V, error)
}
