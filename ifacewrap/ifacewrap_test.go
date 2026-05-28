package ifacewrap

import (
	"testing"

	"github.com/gontract/gontract"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	type input struct {
		a int
		b int
	}

	type output struct {
		sum int
	}

	var order []string

	got := Call(
		input{a: 2, b: 3},
		func() output {
			order = append(order, "body")
			return output{sum: 5}
		},
		Contract[input, output]{
			Requirements: []Requirement[input]{
				{
					Predicate: func(in input) bool {
						order = append(order, "require")
						return in.a >= 0
					},
					Message: "a must be non-negative",
				},
			},
			Assurances: []Assurance[input, output]{
				{
					Predicate: func(in input, out output) bool {
						order = append(order, "ensure")
						return out.sum == in.a+in.b
					},
					Message: "sum must match operands",
				},
			},
		},
	)

	assert.Equal(t, output{sum: 5}, got)
	assert.Equal(t, []string{"require", "body", "ensure"}, order)
}

func TestCallViolation(t *testing.T) {
	type input struct {
		value int
	}

	type output struct {
		value int
	}

	ret := "ok"
	defer gontract.CatchViolation(&ret)

	Call(
		input{value: -1},
		func() output {
			return output{value: -1}
		},
		Contract[input, output]{
			Requirements: []Requirement[input]{
				{
					Predicate: func(in input) bool {
						return in.value >= 0
					},
					Message: "value must be non-negative",
				},
			},
		},
	)

	assert.Equal(t, "assertion error: requirement not satisfied (value must be non-negative) - software bug in caller!", ret)
}

func TestDo(t *testing.T) {
	type input struct {
		counter *int
	}

	trace := []string{}
	counter := 0

	Do(
		input{counter: &counter},
		func() {
			trace = append(trace, "body")
			counter++
		},
		Action[input]{
			Requirements: []Requirement[input]{
				{
					Predicate: func(in input) bool {
						trace = append(trace, "require")
						return in.counter != nil
					},
					Message: "counter must not be nil",
				},
			},
			Assurances: []Requirement[input]{
				{
					Predicate: func(in input) bool {
						trace = append(trace, "ensure")
						return *in.counter == 1
					},
					Message: "counter must be incremented",
				},
			},
		},
	)

	assert.Equal(t, 1, counter)
	assert.Equal(t, []string{"require", "body", "ensure"}, trace)
}

func TestCallLegacyHooks(t *testing.T) {
	type input struct {
		value int
	}

	type output struct {
		value int
	}

	var order []string

	got := Call(
		input{value: 3},
		func() output {
			order = append(order, "body")
			return output{value: 3}
		},
		Contract[input, output]{
			Require: func(in input) {
				order = append(order, "require")
				gontract.Require(in.value > 0, "value must be positive")
			},
			Ensure: func(in input, out output) {
				order = append(order, "ensure")
				gontract.Ensure(out.value == in.value, "value must round-trip")
			},
		},
	)

	assert.Equal(t, output{value: 3}, got)
	assert.Equal(t, []string{"require", "body", "ensure"}, order)
}
