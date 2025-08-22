package common

type Configurable[T any] interface {
	SetConfig(config T) error
	Config() (T, error)
}
