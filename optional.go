package optional

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Optional is a container that may or may not contains a non-nil value.
type Optional[T comparable] struct {
	val *T
}

// New returns an Optional with the value if it's not nil, otherwise returns and empty Optional.
func New[T comparable](v *T) *Optional[T] {
	opt := new(Optional[T])
	opt.val = v
	return opt
}

// Of returns an Optional instance with the specified value.
func Of[T comparable](v T) *Optional[T] {
	opt := new(Optional[T])
	opt.val = &v
	return opt
}

// Empty returns an empty Optional instance.
func Empty[T comparable]() *Optional[T] {
	opt := new(Optional[T])
	return opt
}

// Equals indicates whether some other value is equals to this Optional.
func (opt *Optional[T]) Equals(other any) bool {
	if opt == other {
		return true
	}

	if v, ok := other.(T); ok {
		if opt.val == nil {
			return false
		} else {
			return *opt.val == v
		}
	}

	ov, ok := other.(*Optional[T])
	if !ok {
		return false
	}

	if opt.val == nil && ov.val == nil {
		return true
	} else if opt.val == nil || ov.val == nil {
		return false
	}

	return *opt.val == *ov.val
}

// Filter returns the Optional if the value is present and matches the given predicate, otherwise
// returns an empty Optional.
func (opt *Optional[T]) Filter(predicate func(v T) bool) *Optional[T] {
	if opt.IsEmpty() {
		return opt
	}

	ok := predicate(*opt.val)
	if ok {
		return opt
	} else {
		return Empty[T]()
	}
}

// Get returns the value if a value is present in the Optional, otherwise returns an
// ErrNoSuchValue.
func (opt *Optional[T]) Get() (T, error) {
	if opt.IsEmpty() {
		var zero T
		return zero, ErrNoSuchValue
	}
	return *opt.val, nil
}

// GetPanic returns the value if a value is present in the Optional, otherwise panic
// ErrNoSuchValue.
func (opt *Optional[T]) GetPanic() T {
	if opt.IsEmpty() {
		panic(ErrNoSuchValue)
	}
	return *opt.val
}

// IfPresent performs the given action with the value if it is present, otherwise does nothing.
func (opt *Optional[T]) IfPresent(action func(v T)) {
	if opt.IsPresent() {
		action(*opt.val)
	}
}

// IfPresentOrElse performs the given action with the value if it is present, otherwise performs
// the given empty-based action.
func (opt *Optional[T]) IfPresentOrElse(action func(v T), emptyAction func()) {
	if opt.IsPresent() {
		action(*opt.val)
	} else {
		emptyAction()
	}
}

// IsEmpty return true if there is not a value present, otherwise false.
func (opt *Optional[T]) IsEmpty() bool {
	return opt.val == nil
}

// IsPresent return true if there is a value present, otherwise false.
func (opt *Optional[T]) IsPresent() bool {
	return opt.val != nil
}

// OrElse returns the value if present, otherwise returns other.
func (opt *Optional[T]) OrElse(other T) T {
	if opt.IsEmpty() {
		return other
	}
	return *opt.val
}

// OrElseGet returns the value if present, otherwise returns the result produces by the supplier
// function.
func (opt *Optional[T]) OrElseGet(supplier func() T) T {
	if opt.IsPresent() {
		return *opt.val
	}
	return supplier()
}

// String returns a string for representing the value.
func (opt *Optional[T]) String() string {
	if opt.IsEmpty() {
		return "Optional[" + reflect.TypeOf((*T)(nil)).Elem().String() + "].Empty"
	}

	rv := reflect.ValueOf(opt.val)
	sf := rv.MethodByName("String")
	if sf.IsValid() && sf.Kind() == reflect.Func {
		sft := sf.Type()
		if sft.NumIn() == 0 && sft.NumOut() == 1 && sf.Type().Out(0).Kind() == reflect.String {
			out := sf.Call(nil)
			return out[0].String()
		}
	}

	return fmt.Sprintf("%v", *opt.val)
}

// MarshalJSON marshals the value and returns into valid JSON.
func (opt *Optional[T]) MarshalJSON() ([]byte, error) {
	if opt.IsEmpty() {
		return []byte("null"), nil
	}

	return json.Marshal(*opt.val)
}

// UnmarshalJSON unmarshal a JSON to an Optional value.
func (opt *Optional[T]) UnmarshalJSON(b []byte) error {
	v := new(T)
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.val = v
	return nil
}
