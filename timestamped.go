package common

type UnixTimestamped interface {
	UnixNano() int64
}
