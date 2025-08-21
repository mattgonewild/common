package common

type Cursor[T UnixTimestamped] interface {
	Cursor() T
	Seek(pointUnixTime int64) error
	Start() error
	Previous() bool
	Next() (bool, error)
	End() error
	Point() int64
}
