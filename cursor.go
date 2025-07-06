package common

import "time"

type Cursor[T Timestamped] interface {
	Cursor() (T, bool)
	Seek(point time.Time) bool

	Start()
	Previous() bool
	Next() bool
	End()
}
