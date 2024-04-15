package optional

// Optional is a container that may or may not contains a non-nil value.
type Optional[T comparable] struct {
	val *T
}

// New returns an Optional instance with the specified value.
func New[T comparable](v T) *Optional[T] {
	opt := new(Optional[T])
	opt.val = &v
	return opt
}

// NewNilable returns an Optional with the value if it's not nil, otherwise returns and empty
// Optional.
func NewNilable[T comparable](v *T) *Optional[T] {
	opt := new(Optional[T])
	opt.val = v
	return opt
}

// Empty returns an empty Optional instance.
func Empty[T comparable]() *Optional[T] {
	opt := new(Optional[T])
	return opt
}

// Equals indicates whether some other value is equals to this Optional.
func (opt *Optional[T]) Equals(other T) bool {
	if !opt.IsPresent() {
		return false
	}

	return *opt.val == other
}

// Get returns the value if a value is present in the Optional, otherwise returns an
// ErrNoSuchValue.
func (opt *Optional[T]) Get() (T, error) {
	if !opt.IsPresent() {
		var zero T
		return zero, ErrNoSuchValue
	}
	return *opt.val, nil
}

// GetPanic returns the value if a value is present in the Optional, otherwise panic
// ErrNoSuchValue.
func (opt *Optional[T]) GetPanic() T {
	if !opt.IsPresent() {
		panic(ErrNoSuchValue)
	}
	return *opt.val
}

// IsPresent return true if there is a value present, otherwise false.
func (opt *Optional[T]) IsPresent() bool {
	return opt.val != nil
}

// OrElse returns the value if present, otherwise returns other
func (opt *Optional[T]) OrElse(other T) T {
	if !opt.IsPresent() {
		return other
	}
	return *opt.val
}
