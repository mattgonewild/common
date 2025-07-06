package common

import "time"

type Log[T Timestamped] interface {
	Append(element T) error
	Tail() (T, bool)
	NewCursor() Cursor[T]
	NewBoundedCursor(through time.Time) Cursor[T]
}
