package common

import "time"

type Cursor[T Timestamped] interface {
	Cursor() T
	Seek(point time.Time) error
	Start() error
	Previous() bool
	Next() (bool, error)
	End() error
	Point() time.Time
}
