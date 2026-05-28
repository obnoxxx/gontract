package ifacewrap

import (
	"fmt"

	"github.com/gontract/gontract"
)

// Requirement describes a precondition for a wrapped method call.
//
// Predicate receives the grouped method input and Message becomes the
// gontract.Require violation text when Predicate returns false.
type Requirement[In any] struct {
	Predicate func(In) bool
	Message   string
}

// Assurance describes a postcondition for a wrapped method call.
//
// Predicate receives the grouped method input and output, and Message becomes
// the gontract.Ensure violation text when Predicate returns false.
type Assurance[In any, Out any] struct {
	Predicate func(In, Out) bool
	Message   string
}

// Contract describes pre- and postconditions for a wrapped method call.
//
// In is typically a small struct that groups the method inputs.
// Out is typically either the method result itself or a small struct that groups
// multiple return values.
//
// Require and Ensure provide legacy callback hooks for callers that need custom
// behavior. Requirements and Assurances are the preferred declarative form and
// are automatically enforced via gontract.Require and gontract.Ensure.
type Contract[In any, Out any] struct {
	Require      func(In)
	Ensure       func(In, Out)
	Requirements []Requirement[In]
	Assurances   []Assurance[In, Out]
}

// Call executes body with the given contract.
//
// Require and Requirements, when present, run before body. Ensure and
// Assurances, when present, run after body and observe the final returned value.
func Call[In any, Out any](in In, body func() Out, contract Contract[In, Out]) (out Out) {
	if contract.Require != nil {
		contract.Require(in)
	}

	runRequirements(in, contract.Requirements)

	out = body()

	if contract.Ensure != nil {
		contract.Ensure(in, out)
	}

	runAssurances(in, out, contract.Assurances)

	return
}

// Action describes pre- and postconditions for a wrapped method that does not
// return a value.
type Action[In any] struct {
	Require      func(In)
	Ensure       func(In)
	Requirements []Requirement[In]
	Assurances   []Requirement[In]
}

// Do executes body with the given action contract.
func Do[In any](in In, body func(), action Action[In]) {
	if action.Require != nil {
		action.Require(in)
	}

	runRequirements(in, action.Requirements)

	body()

	if action.Ensure != nil {
		action.Ensure(in)
	}

	runRequirements(in, action.Assurances)
}

func runRequirements[In any](in In, requirements []Requirement[In]) {
	for i, requirement := range requirements {
		if requirement.Predicate == nil {
			panic(fmt.Sprintf("ifacewrap: requirement %d has nil predicate", i))
		}

		gontract.Require(requirement.Predicate(in), requirement.Message)
	}
}

func runAssurances[In any, Out any](in In, out Out, assurances []Assurance[In, Out]) {
	for i, assurance := range assurances {
		if assurance.Predicate == nil {
			panic(fmt.Sprintf("ifacewrap: assurance %d has nil predicate", i))
		}

		gontract.Ensure(assurance.Predicate(in, out), assurance.Message)
	}
}
