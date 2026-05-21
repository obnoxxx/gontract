include meta.mk

include go.mk
include examples.mk

.DEFAULT_GOAL := all


.PHONY: all
all: build.examples

.PHONY: lint
lint: lint.go

.PHONY: test
test: test.go

.PHONY: check
check: all check.build.all test run.examples.good

.PHONY: clean
clean: clean.examples
