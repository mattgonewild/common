package common

import "time"

type Log[T Timestamped] interface {
	Append(element T) error
	Tail() (T, bool)
	NewCursor(start time.Time) Cursor[T]
	NewBoundedCursor(start, through time.Time) Cursor[T]
}
