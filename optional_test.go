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

func TestOptionalFilter(t *testing.T) {
	a := assert.New(t)
	predicate := func(s string) bool { return s == "Hello world" }
	empty := optional.Empty[string]()

	opt := optional.Empty[string]()
	ret := opt.Filter(predicate)
	a.DeepEqualNow(ret, empty)

	opt = optional.New("test")
	ret = opt.Filter(predicate)
	a.DeepEqualNow(ret, empty)

	opt = optional.New("Hello world")
	ret = opt.Filter(predicate)
	a.DeepEqualNow(ret, opt)
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

func TestOptionalIfPresent(t *testing.T) {
	a := assert.New(t)
	run := false

	opt := optional.Empty[string]()
	opt.IfPresent(func(v string) {
		a.FailNow()
	})
	a.NotTrueNow(run)

	opt = optional.New("Hello world")
	opt.IfPresent(func(v string) {
		a.EqualNow(v, "Hello world")
		run = true
	})
	a.TrueNow(run)
}

func TestOptionalIfPresentOrElse(t *testing.T) {
	a := assert.New(t)
	actionRun := false
	emptyRun := false

	opt := optional.Empty[string]()
	opt.IfPresentOrElse(func(v string) {
		a.FailNow()
	}, func() {
		emptyRun = true
	})
	a.NotTrueNow(actionRun)
	a.TrueNow(emptyRun)

	actionRun = false
	emptyRun = false
	opt = optional.New("Hello world")
	opt.IfPresentOrElse(func(v string) {
		a.EqualNow(v, "Hello world")
		actionRun = true
	}, func() {
		a.FailNow()
	})
	a.TrueNow(actionRun)
	a.NotTrueNow(emptyRun)
}

func TestOptionalIsPresent(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.NotTrueNow(opt.IsPresent())

	opt = optional.New("Hello world")
	a.TrueNow(opt.IsPresent())
}

func TestOptionalIsEmpty(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.TrueNow(opt.IsEmpty())

	opt = optional.New("Hello world")
	a.NotTrueNow(opt.IsEmpty())
}

func TestOptionalOrElse(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.EqualNow(opt.OrElse("default"), "default")

	opt = optional.New("Hello world")
	a.EqualNow(opt.OrElse("default"), "Hello world")
}

func TestOptionalOrElseGet(t *testing.T) {
	a := assert.New(t)
	action := func() string {
		return "default"
	}

	opt := optional.Empty[string]()
	a.EqualNow(opt.OrElseGet(action), "default")

	opt = optional.New("Hello world")
	a.EqualNow(opt.OrElseGet(action), "Hello world")
}
