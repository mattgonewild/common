package common

type Comparable interface {
	Compare(comparable Comparable) int
}
