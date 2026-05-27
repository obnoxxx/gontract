CHECKMAKE := go run github.com/checkmake/checkmake/cmd/checkmake@v0.3.2
GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint.go.vet
lint.go.vet:
	@echo "vetting go code..."
	@go vet ./...

.PHONY: lint.go.fmt
lint.go.fmt:
	@echo "Checking go formatting..."
	@if [ -n "$$(gofmt -l .)" ]; then \
			echo "Files need formatting:"; \
				gofmt -l .; \
				exit 1; \
	else \
		echo "All files formatted correctly."; \
	fi

.PHONY: lint.make
lint.make: Makefile
	@echo "checking the Makefile..."
	@$(CHECKMAKE) Makefile
	@echo "The Makefile is OK."

.PHONY: fix.go.fmt
fix.go.fmt: # fix go formatting (if needed)
	@go fmt ./...

.PHONY: test.go
test.go: lint.go
	@go test ./...

.PHONY: check.build.all
check.build.all: lint.go
	@echo "building go code ..."
	@go build ./...

.PHONY: golangci-lint
golangci-lint:
	@echo "linting go code ..."
	@$(GOLANGCI_LINT) run

.PHONY: lint.go
lint.go: lint.go.fmt lint.go.vet golangci-lint
