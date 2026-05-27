package ifacewrap

// Contract describes pre- and postconditions for a wrapped method call.
//
// In is typically a small struct that groups the method inputs.
// Out is typically either the method result itself or a small struct that groups
// multiple return values.
type Contract[In any, Out any] struct {
	Require func(In)
	Ensure  func(In, Out)
}

// Call executes body with the given contract.
//
// Require, when present, runs before body. Ensure, when present, runs after body
// and observes the final returned value.
func Call[In any, Out any](in In, body func() Out, contract Contract[In, Out]) (out Out) {
	if contract.Require != nil {
		contract.Require(in)
	}

	out = body()

	if contract.Ensure != nil {
		contract.Ensure(in, out)
	}

	return
}

// Action describes pre- and postconditions for a wrapped method that does not
// return a value.
type Action[In any] struct {
	Require func(In)
	Ensure  func(In)
}

// Do executes body with the given action contract.
func Do[In any](in In, body func(), action Action[In]) {
	if action.Require != nil {
		action.Require(in)
	}

	body()

	if action.Ensure != nil {
		action.Ensure(in)
	}
}
