
[![Build Status](https://github.com/gontract/gontract/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/gontract/gontract/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gontract/gontract)](https://goreportcard.com/report/github.com/gontract/gontract)
[![GPL license](https://img.shields.io/badge/license-LGPL-blue.svg)](http://opensource.org/licenses/LGPL)
[![Go Reference](https://pkg.go.dev/badge/github.com/gontract/gontract.svg)](https://pkg.go.dev/github.com/gontract/gontract)

# gontract
**gontract** is a require-ensure-library for Go that provides a contract-programming framework. To enable Design by Contract (DbC), it aims to bridge the "procedural gap" in Go by providing explicit `Require` and `Ensure` checks for preconditions and postconditions in Go code.

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

A very natural pattern is to put posconditions into a defer statement like so:

```go

func myfunc(args...) {

// postcondition(s):
defer func() {

Ensure(..)
..Ensure(..)

}()


// precondtions:
Require(..)
...
Require(..)
...
//implementation
...

```

## Interface wrappers

For interface-based designs, keep the interface itself plain and apply
preconditions and postconditions in a wrapper implementation around the real
implementation.

The companion package `github.com/gontract/gontract/ifacewrap` provides small
helpers for that style. See `cmd/examples/ifacewrap_division` for a minimal
example.














