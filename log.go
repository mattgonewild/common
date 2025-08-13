package common

import "time"

type Log[T Timestamped] interface {
	Append(element T) error
	Tail() (T, bool)

	NewCursorBefore(end time.Time) (Cursor[T], error)
	NewCursorAt(floor time.Time) (Cursor[T], error)
	NewCursorAfter(start time.Time) (Cursor[T], error)
}
