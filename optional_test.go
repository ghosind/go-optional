package optional_test

import (
	"encoding/json"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-optional"
)

type TestStruct struct {
	Val string
}

func (v *TestStruct) String() string {
	return "TestStruct<" + v.Val + ">"
}

type TestStr string

func (s TestStr) String() string {
	return "TestStr<" + string(s) + ">"
}

func TestOptional(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.NotTrueNow(opt.IsPresent())
	a.EqualNow(opt.OrElse("default"), "default")

	opt = optional.Of("Hello world")
	a.TrueNow(opt.IsPresent())
	a.EqualNow(opt.OrElse("default"), "Hello world")

	var sp *string
	opt = optional.New(sp)
	a.NotTrueNow(opt.IsPresent())
	a.EqualNow(opt.OrElse("default"), "default")

	s := "some text"
	sp = &s
	opt = optional.New(sp)
	a.TrueNow(opt.IsPresent())
	a.EqualNow(opt.OrElse("default"), "some text")
}

func TestOptionalEquals(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.NotTrueNow(opt.Equals("not equals"))
	a.NotTrueNow(opt.Equals(optional.Empty[int]()))
	a.TrueNow(opt.Equals(optional.Empty[string]()))
	a.NotTrueNow(opt.Equals(optional.Of("not equals")))

	opt = optional.Of("Hello world")
	a.NotTrueNow(opt.Equals("not equals"))
	a.TrueNow(opt.Equals("Hello world"))
	a.NotTrueNow(opt.Equals(optional.Empty[string]()))
	a.NotTrueNow(opt.Equals(optional.Of("not equals")))
	a.TrueNow(opt.Equals(optional.Of("Hello world")))
	a.TrueNow(opt.Equals(opt))
}

func TestOptionalFilter(t *testing.T) {
	a := assert.New(t)
	predicate := func(s string) bool { return s == "Hello world" }
	empty := optional.Empty[string]()

	opt := optional.Empty[string]()
	ret := opt.Filter(predicate)
	a.DeepEqualNow(ret, empty)

	opt = optional.Of("test")
	ret = opt.Filter(predicate)
	a.DeepEqualNow(ret, empty)

	opt = optional.Of("Hello world")
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

	opt = optional.Of("Hello world")

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

	opt = optional.Of("Hello world")
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
	opt = optional.Of("Hello world")
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

	opt = optional.Of("Hello world")
	a.TrueNow(opt.IsPresent())
}

func TestOptionalIsEmpty(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.TrueNow(opt.IsEmpty())

	opt = optional.Of("Hello world")
	a.NotTrueNow(opt.IsEmpty())
}

func TestOptionalOrElse(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	a.EqualNow(opt.OrElse("default"), "default")

	opt = optional.Of("Hello world")
	a.EqualNow(opt.OrElse("default"), "Hello world")
}

func TestOptionalOrElseGet(t *testing.T) {
	a := assert.New(t)
	action := func() string {
		return "default"
	}

	opt := optional.Empty[string]()
	a.EqualNow(opt.OrElseGet(action), "default")

	opt = optional.Of("Hello world")
	a.EqualNow(opt.OrElseGet(action), "Hello world")
}

func TestString(t *testing.T) {
	a := assert.New(t)

	intOpt := optional.Empty[int]()
	a.EqualNow(intOpt.String(), "Optional[int].Empty")
	intOpt = optional.Of(1)
	a.EqualNow(intOpt.String(), "1")

	strOpt := optional.Empty[string]()
	a.EqualNow(strOpt.String(), "Optional[string].Empty")
	strOpt = optional.Of("test")
	a.EqualNow(strOpt.String(), "test")

	valOpt := optional.Empty[TestStruct]()
	a.EqualNow(valOpt.String(), "Optional[optional_test.TestStruct].Empty")
	valOpt = optional.Of(TestStruct{Val: "Hello world"})
	a.EqualNow(valOpt.String(), "TestStruct<Hello world>")

	vpOpt := optional.Empty[*TestStruct]()
	a.EqualNow(vpOpt.String(), "Optional[*optional_test.TestStruct].Empty")
	vpOpt = optional.Of(&TestStruct{Val: "Hello world"})
	a.EqualNow(vpOpt.String(), "TestStruct<Hello world>")

	tsOpt := optional.Empty[TestStr]()
	a.EqualNow(tsOpt.String(), "Optional[optional_test.TestStr].Empty")
	tsOpt = optional.Of(TestStr("Hello world"))
	a.EqualNow(tsOpt.String(), "TestStr<Hello world>")
}

func TestOptionalMarshalJSON(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	text, err := json.Marshal(opt)
	a.NilNow(err)
	a.EqualNow(string(text), `null`)

	text, err = opt.MarshalJSON()
	a.NilNow(err)
	a.EqualNow(string(text), `null`)

	opt = optional.Of("Hello world")
	text, err = json.Marshal(opt)
	a.NilNow(err)
	a.EqualNow(string(text), `"Hello world"`)

	text, err = opt.MarshalJSON()
	a.NilNow(err)
	a.EqualNow(string(text), `"Hello world"`)

	sv := new(struct {
		Val *optional.Optional[string] `json:"v"`
	})
	text, err = json.Marshal(sv)
	a.NilNow(err)
	a.EqualNow(string(text), `{"v":null}`)

	sv.Val = optional.Of("Hello world")
	text, err = json.Marshal(sv)
	a.NilNow(err)
	a.EqualNow(string(text), `{"v":"Hello world"}`)
}

func TestOptionalUnmarshalJSON(t *testing.T) {
	a := assert.New(t)

	opt := optional.Empty[string]()
	err := json.Unmarshal([]byte(`null`), opt)
	a.NilNow(err)
	a.TrueNow(opt.IsEmpty())

	err = json.Unmarshal([]byte(`"Hello world"`), opt)
	a.NilNow(err)
	a.TrueNow(opt.IsPresent())
	a.EqualNow(opt.GetPanic(), "Hello world")

	v := new(struct {
		Val optional.Optional[string] `json:"v"`
	})
	err = json.Unmarshal([]byte(`{"v":null}`), v)
	a.NilNow(err)
	a.TrueNow(v.Val.IsEmpty())

	err = json.Unmarshal([]byte(`{"v":"Hello world"}`), v)
	a.NilNow(err)
	a.TrueNow(v.Val.IsPresent())
	a.EqualNow(v.Val.GetPanic(), "Hello world")

	pv := new(struct {
		Val *optional.Optional[string] `json:"v"`
	})
	err = json.Unmarshal([]byte(`{"v":null}`), pv)
	a.NilNow(err)
	a.NilNow(pv.Val)

	err = json.Unmarshal([]byte(`{"v":"Hello world"}`), pv)
	a.NilNow(err)
	a.TrueNow(pv.Val.IsPresent())
	a.EqualNow(pv.Val.GetPanic(), "Hello world")

	err = opt.UnmarshalJSON([]byte(`null`))
	a.NilNow(err)
	a.TrueNow(opt.IsEmpty())

	err = opt.UnmarshalJSON([]byte(`"Hello world"`))
	a.NilNow(err)
	a.TrueNow(opt.IsPresent())
	a.EqualNow(opt.GetPanic(), "Hello world")

	err = opt.UnmarshalJSON([]byte(`unknown`))
	a.NotNilNow(err)
}
