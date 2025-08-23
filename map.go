package common

import "iter"

type Map[V any] interface {
	Set(key Key, value V) error
	Get(key Key) (V, error)
	Delete(key Key) error
}

type SortedMap[K Comparable[K], V any] interface {
	Set(key K, value V) error
	Get(key K) (V, error)
	Delete(key K) error
	All() iter.Seq2[K, V]
}
