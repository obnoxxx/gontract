
[![Build Status](https://github.com/gontract/gontract/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/gontract/gontract/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gontract/gontract)](https://goreportcard.com/report/github.com/gontract/gontract)
[![GPL license](https://img.shields.io/badge/license-LGPL-blue.svg)](http://opensource.org/licenses/LGPL)
[![Go Reference](https://pkg.go.dev/badge/github.com/gontract/gontract.svg)](https://pkg.go.dev/github.com/gontract/gontract)

# gontract
**gontract** is a **require-ensure-library** for Go that provides a contract-programming framework. To enable partial Design by Contract (DbC), it aims to bridge the "procedural gap" in Go by providing explicit `Require` and `Ensure` checks for writing preconditions and postconditions naturally and idiomatically  in Go .

## Core Pillars

* **Contract Decoupling:** Separating safety logic from business logic.
* **Safety Assertions:** Provides immediate feedback upon contract violation during development.

---

## API Reference

| Function | Purpose | Usage |
| :--- | :--- | :--- |
| `Require(predicate bool, msg string)` | **Precondition**: Verifies caller-provided input. | Start of function. |
| `Ensure(predicate bool, msg string)` | **Postcondition**: Verifies function logic/return values. | Before return. |

---


## Best Practice

The typical use of gontract is to wrap the implementation body of a function in a
Require-Ensure-sandwich. Here is an incomplete, non-working example to illustrate the idea:

```go
func myfunc(args) {
// precondition(s):
Require(predicate, "message")
...
Require(predicate, "message")
// implementation
...
// postcondition(s):
Ensure(predicate, "message")
Ensure(predicate, "message")
...
}
```


A very natural pattern is to put postconditions into a defer statement like so:

```go

func myfunc(args) {
// precondition(s):
Require(predicate, "message")
...
Require(predicate, "message")
...

// postcondition(s):
defer func() {

Ensure(predicate, "message")
...
Ensure(predicate, "message")

}()


// implementation
...
}
```

These examples are not intended to work, just to illustrate ideas.

For real, working examples, look into [the examples folder](cmd/examples/)


## Interface wrappers (ifacewrap)

For interface-based designs, it is best to keep the interface itself plain and apply
preconditions and postconditions in a wrapper implementation around the real
implementation instead of clobbering the function implementations with the contract implementations using
`Require` and `Ensure`.

The companion package `github.com/gontract/gontract/ifacewrap` provides small
helpers for that style. Its declarative `Requirements` and `Assurances` run
through `gontract.Require` and `gontract.Ensure` automatically, so wrapper code
only needs to provide predicates and messages. See
`cmd/examples/ifacewrap_division` for a minimal example.

## function wrappers (funcwrap)

In addition to ifacewrap, the companion package `github.com/gontract/gontract/funcwrap`
provides helpers for defining a function along with its contract.
See `cmd/examples/funcwrap_division` for a simple example.

## Example and split-out: manc

A small math helper library has been split out of the division examples into a separate project [manc](https://github.com/obnoxxx/manc).

To increase the cohesion but hopefully not the confusion, manc uses gontract for function contracts while being used in contract specifications in the division examples in the gontract projects itself.
