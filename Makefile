

GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest


.PHONY: lint.go.vet
lint.go.vet:
	@echo "vetting go code..."
	@go vet ./...


.PHOHY: lint.go.fmt
lint.go.fmt:
	@echo "Checking go formatting..."
	@if [ -n "$$(gofmt -l .)" ]; then \
			echo "Files need formatting:"; \
				gofmt -l .; \
				exit 1; \
	else \
		echo "All files formatted correctly."; \
	fi

.PHONY: fix.go.fmt
fix.go.fmt: #␣fix␣go␣formatting (if needed)
	@ go fmt ./...

.PHONY: test
test: lint
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
check: test
