TOP_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
SHELL = /bin/bash

test:
	@echo "### Running unit tests..."
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic ./internal/... ./cmd/...
