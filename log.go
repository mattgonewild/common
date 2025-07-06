package common

type Log[T Timestamped] interface {
	Append(element T) error
	Tail() (T, bool)
	NewCursor() Cursor[T]
}
