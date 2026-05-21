# Copilot Instructions

## Build, test, and lint commands

- Use **Go 1.26**. CI installs Go 1.26 and the module declares `go 1.26`.
- Run the full repository check with `make check`.
- Run linting with `make lint`.
- Run tests with `make test` or `go test ./...`.
- Run a single test by name with `go test ./... -run '^TestRequire$'`.
- Run a single package's tests with `go test ./cmd/examples/division`.
- Run a single test in one package with `go test ./cmd/examples/division -run '^TestDivide$'`.
- Build all example binaries with `make build.examples` (alias: `make examples.demos`).
- Run the non-negative examples with `make run.examples.valid` (alias: `make run.examples.non-negative`).
- Apply Go formatting with `make fix.go.fmt`.
- If you need a build smoke test for all packages and examples, use `go build ./...`.

## High-level architecture

- The root `gontract` package is the library API. Its public surface is intentionally small: `Require`, `Ensure`, `Condition`, `CatchViolation`, and the assertion toggles `EnableAssertions`, `DisableAssertions`, and `AssertionsAreEnabled`.
- `Require` and `Ensure` are thin wrappers around `Condition`. `Condition` is where violation messages are assembled and where failed contracts are forwarded to `gitlab.com/stone.code/assert`.
- Contract failures distinguish **caller bugs** from **callee bugs** through `Kind`: `KindRequire` reports a broken precondition by the caller, and `KindEnsure` reports a broken postcondition by the callee.
- Assertion behavior is controlled by the package-global `assertionsEnabled` flag. Assertions are enabled by default; when disabled, `Condition` still returns the formatted message but skips the panic path.
- `cmd/examples/*` is part of the repository's design, not just demo code. Those example programs exercise the library in realistic success and failure scenarios, and `go test ./...` includes them alongside the root package tests.

## Key conventions

- Follow the design-by-contract model described in `README.md`: use `Require` for preconditions at function entry and `Ensure` for postconditions owned by the implementation.
- Postconditions are often written with named return values and `defer`, so the `Ensure` call observes the final result after the function body has assigned it.
- The exact wording of violation messages is test-covered in the root package. If contract message text changes, update tests that assert the full recovered string.
- Tests for contract violations recover from panics instead of expecting ordinary errors. In the root package, wrappers defer `CatchViolation(&ret)` to capture the panic text into a string return value.
- Example-package tests commonly use local `catch_violation` helpers that recover panics into sentinel numeric values like `-1.0`. Those examples intentionally include both success cases and contract-breaking cases; do not "fix" the negative examples unless the intended demonstration changes.
