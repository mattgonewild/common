package common

import "time"

type Log[T Timestamped] interface {
	Append(element T) error
	Tail() (T, bool)
	NewCursor(anchor time.Time, end bool) (Cursor[T], error)
}
