#!/usr/bin/env bash


# Exit immediately if a command exits with a non‑zero status.
set -e

# Format Go code
go fmt ./...

# Does this code compile?
go build -o simctl ./cmd/simctl

# how do we see if test files exist for each file?

# Run tests for the current module and all its sub‑packages -- Do the tests pass?
go test -v ./...

# Run tests with race detector, detects concurrency bugs
go test -race -v ./...

# Assess test coverage
go test -cover ./...

# Check for static analysis issues "Does anything look suspicious?"
go vet -v ./...

# Are there bugs, bad practices or code smells?
staticcheck ./...
