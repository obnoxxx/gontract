EXAMPLES_DIR := $(ROOT_DIR)/cmd/examples
EXAMPLES_BUILD_DIR := $(ROOT_DIR)/.build/examples

EXAMPLE_NAMES := $(sort $(notdir $(wildcard $(EXAMPLES_DIR)/*)))
NEGATIVE_EXAMPLE_NAMES := sqrt_fail_post sqrt_fail_pre

EXAMPLES_VALID := $(filter-out $(NEGATIVE_EXAMPLE_NAMES),$(EXAMPLE_NAMES))
EXAMPLES_DEMO_BINS := $(addprefix $(EXAMPLES_BUILD_DIR)/,$(EXAMPLE_NAMES))

$(EXAMPLES_BUILD_DIR):
	@mkdir -p $@

$(EXAMPLES_BUILD_DIR)/%: | $(EXAMPLES_BUILD_DIR)
	@go build -o $@ ./cmd/examples/$*

.PHONY: build.examples
build.examples: $(EXAMPLES_DEMO_BINS)

.PHONY: examples.demos
examples.demos: build.examples

.PHONY: run.examples.valid
run.examples.valid:
	@for example in $(EXAMPLES_VALID); do \
		echo "running $$example..."; \
		go run ./cmd/examples/$$example; \
	done

.PHONY: run.examples.non-negative
run.examples.non-negative: run.examples.valid

.PHONY: run.examples.success
run.examples.success: run.examples.valid

.PHONY: run.examples.good
run.examples.good: run.examples.valid

.PHONY: clean.examples
clean.examples:
	@rm -rf $(EXAMPLES_BUILD_DIR)
