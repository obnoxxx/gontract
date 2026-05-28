// Package ifacewrap provides small helpers for writing contract-aware wrappers
// around Go interfaces.
//
// The package does not attempt to generate wrappers or proxy arbitrary
// interfaces. Instead, it keeps wrapper methods explicit and idiomatic while
// factoring out the repetitive "require/body/ensure" control flow while
// automatically dispatching declarative checks through gontract.
package ifacewrap
