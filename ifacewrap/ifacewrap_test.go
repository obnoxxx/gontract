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
			Require: func(in input) {
				order = append(order, "require")
				gontract.Require(in.a >= 0, "a must be non-negative")
			},
			Ensure: func(in input, out output) {
				order = append(order, "ensure")
				gontract.Ensure(out.sum == in.a+in.b, "sum must match operands")
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
			Require: func(in input) {
				gontract.Require(in.value >= 0, "value must be non-negative")
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
			Require: func(in input) {
				trace = append(trace, "require")
				gontract.Require(in.counter != nil, "counter must not be nil")
			},
			Ensure: func(in input) {
				trace = append(trace, "ensure")
				gontract.Ensure(*in.counter == 1, "counter must be incremented")
			},
		},
	)

	assert.Equal(t, 1, counter)
	assert.Equal(t, []string{"require", "body", "ensure"}, trace)
}
