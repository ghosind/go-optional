package optional

import "errors"

var (
	ErrNilFunction = errors.New("nil function")
	ErrNoSuchValue = errors.New("no such value")
)
