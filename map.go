package common

import "iter"

type Map[T any] interface {
	Set(key Key, value T) error
	Get(key Key) (T, error)
	Delete(key Key) error
}

type SortedMap[T Comparable[T]] interface {
	Set(T) error
	Get(T) (T, error)
	Delete(T) error
	All() iter.Seq[T]
}
