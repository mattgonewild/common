package common

import "time"

type Cursor[T Timestamped] interface {
	Cursor() (T, bool)
	Seek(point time.Time) bool

	End()
	Previous() bool
	Next() bool
	Begin()
}
