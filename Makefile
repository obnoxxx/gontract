include meta.mk

include go.mk
include examples.mk

.DEFAULT_GOAL := all


.PHONY: all
all: build.examples

.PHONY: lint
lint: lint.go lint.make

.PHONY: test
test: test.go

.PHONY: check
check: all lint check.build.all test run.examples.good

.PHONY: clean
clean: clean.examples
