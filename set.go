package common

import "iter"

type Set[T Comparable[T]] interface {
	Put(element T) error
	Get(element T) (T, error)
	Delete(element T) error
	All() iter.Seq[T]
}
