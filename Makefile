# Makefile for zpgxpool
.PHONY: all test examples docs fmt vet ci clean

GOCMD ?= go
GOTEST ?= $(GOCMD) test
GOMOD ?= $(GOCMD) mod

all: test

# Run all unit tests in the repo
test:
	$(GOTEST) ./...

# Run only examples tests
examples:
	cd examples && $(GOTEST) ./...

# Format code
fmt:
	$(GOCMD) fmt ./...

# Run vet/lint
vet:
	$(GOCMD) vet ./...

# Run CI locally (format + vet + tests)
ci: fmt vet test examples

# Clean module cache (useful for CI debugging)
clean:
	$(GOCMD) clean -modcache