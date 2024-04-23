# go-optional

![test](https://github.com/ghosind/go-optional/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/go-optional)](https://goreportcard.com/report/github.com/ghosind/go-optional)
[![codecov](https://codecov.io/gh/ghosind/go-optional/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-optional)
![Version Badge](https://img.shields.io/github/v/release/ghosind/go-optional)
![License Badge](https://img.shields.io/github/license/ghosind/go-optional)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/go-optional.svg)](https://pkg.go.dev/github.com/ghosind/go-optional)

A container object to describe the specified type value that may or may not contain a non-nil value.

## Installation

You can install the package by the following command.

```sh
go get -u github.com/ghosind/go-optional
```

## Getting Started

You can create an `Optional` instance with a pointer, and perform actions with the `Optional` instance.

```go
s := "Hello world"
vp := &s

val = optional.New(vp)
a.IsPresent() // true
val.OrElse("default string") // Hello world
val.IfPresent(func (s string) {
  fmt.Println(s)
})
// Hello world

// Optional instance with nil value
vp = nil
val := optional.New(vp)
vp.IsPresent() // false
val.OrElse("default string") // default string
val.IfPresent(func (s string) {
  fmt.Println(s)
}) // Not invoked
```

The `Optional` type also supports marshaling the value to a JSON string, or unmarshal from a JSON string.

```go
type Example struct {
  Val *optional.Optional[string] `json:"val"`
}

example := Example{
  Val: optional.Of("Hello world"),
}
out, err := json.Marshal(example)
fmt.Println(string(out))
// {"val":"Hello world"}

json.Unmarshal([]byte("Test"), &example)
fmt.Println(example)
// {Hello world}
```
