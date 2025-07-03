package common

type Key interface {
	Byte() []byte
	Equal(key Key) bool
}
