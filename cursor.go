package common

import "time"

type Cursor[T Timestamped] interface {
	Cursor() (T, error)
	Seek(point time.Time) bool

	Start()
	Previous() bool
	Next() (bool, error)
	End()

	Point() (time.Time, error)
}
