package optional

import "errors"

var (
	// ErrNilFunction indicates the supplier function is nil.
	ErrNilFunction = errors.New("nil function")
	// ErrNoSuchValue indicates the Optional does not contain any value.
	ErrNoSuchValue = errors.New("no such value")
)
