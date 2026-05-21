
MK_SOURCE := Makefile


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
lint.make: $(MK_SOURCE)
	@$(CHECKMAKE) $(MK_SOURCE)

.PHONY: fix.go.fmt
fix.go.fmt: # fix go formatting (if needed)
	@ go fmt ./...

.PHONY: test
test: lint.go
	@go test ./...

.PHONY: golangci-lint
golangci-lint:
	@echo "linting go code ..."
	@$(GOLANGCI_LINT) run

.PHONY: lint.go
lint.go: lint.go.fmt lint.go.vet golangci-lint


.PHONY: lint
lint: lint.go

.PHONY: check
check: lint test
