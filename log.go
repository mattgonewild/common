package common

type Log[T UnixTimestamped] interface {
	Append(element T) error
	Tail() (T, bool)

	NewCursorBefore(endUnixTime int64) (Cursor[T], error)
	NewCursorAt(floorUnixTime int64) (Cursor[T], error)
	NewCursorAfter(startUnixTime int64) (Cursor[T], error)
}
