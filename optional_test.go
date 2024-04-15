package optional_test

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-optional"
)

func TestOptional(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.NotTrueNow(opt.IsPresent())
	a.EqualNow(opt.OrElse("default"), "default")

	opt = optional.New("Hello world")
	a.TrueNow(opt.IsPresent())
	a.EqualNow(opt.OrElse("default"), "Hello world")

	var sp *string
	opt = optional.NewNilable(sp)
	a.NotTrueNow(opt.IsPresent())
	a.EqualNow(opt.OrElse("default"), "default")

	s := "some text"
	sp = &s
	opt = optional.NewNilable(sp)
	a.TrueNow(opt.IsPresent())
	a.EqualNow(opt.OrElse("default"), "some text")
}

func TestOptionalEquals(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.NotTrueNow(opt.Equals("not equals"))

	opt = optional.New("Hello world")
	a.NotTrueNow(opt.Equals("not equals"))
	a.TrueNow(opt.Equals("Hello world"))
}

func TestOptionalGet(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()

	v, err := opt.Get()
	a.EqualNow(err, optional.ErrNoSuchValue)
	a.EqualNow(v, "")

	a.PanicOfNow(func() {
		opt.GetPanic()
		a.FailNow()
	}, optional.ErrNoSuchValue)

	opt = optional.New("Hello world")

	v, err = opt.Get()
	a.NilNow(err)
	a.EqualNow(v, "Hello world")

	a.NotPanicNow(func() {
		v := opt.GetPanic()
		a.EqualNow(v, "Hello world")
	})
}

func TestOptionalIsPresent(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.NotTrueNow(opt.IsPresent())

	opt = optional.New("Hello world")
	a.TrueNow(opt.IsPresent())
}

func TestOptionalOrElse(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.EqualNow(opt.OrElse("default"), "default")

	opt = optional.New("Hello world")
	a.EqualNow(opt.OrElse("default"), "Hello world")
}
