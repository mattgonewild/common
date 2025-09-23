package common

type Registry[K comparable, V any] interface {
	Get(key K) (V, error)
	Release(key K) bool
	Claim(key K) bool
	ClaimOrRegister(key K, new func() V) (V, bool)
}
