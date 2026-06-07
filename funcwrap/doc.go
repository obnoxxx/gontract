// Package funcwrap provides small helpers for writing contract-aware wrappers
// around individual Go functions.
//
// The package keeps wrappers explicit and idiomatic while factoring out the
// repetitive "require/body/ensure" control flow and automatically dispatching
// declarative checks through gontract.
//
// Use Call for functions that return a value and Do for functions that do not.
// A typical wrapper groups the function inputs into a small struct so that
// preconditions and postconditions can inspect the full call state.
//
// Requirements and Assurances are the preferred API. The lower-level
// Require and Ensure callback hooks remain available for contracts that need
// imperative logic instead of a single predicate/message pair.
package funcwrap
