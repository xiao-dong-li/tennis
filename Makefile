# Project configuration
BINARY_NAME := tennis
BUILD_DIR := build
MAIN_PACKAGE := .

# Build flags
GO_FLAGS := -ldflags "-s -w"

# Default target
.PHONY: all
all: clean build-linux build-windows

# Install dependencies
.PHONY: deps
deps:
	@echo ">> Installing dependencies..."
	go mod tidy

# Format source code
.PHONY: fmt
fmt:
	@echo ">> Formatting code..."
	go fmt ./...

# Lint the project (requires golangci-lint)
.PHONY: lint
lint:
	@echo ">> Running linter..."
	golangci-lint run

# Run tests
.PHONY: test
test:
	@echo ">> Running tests..."
	go test -v ./...

# Build for Linux (x86_64)
.PHONY: build-linux
build-linux:
	@echo ">> Building for Linux amd64..."
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_FLAGS) $(MAIN_PACKAGE)

# Build for Windows (x86_64)
build-windows:
	@echo ">> Building for Windows amd64..."
	mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME).exe $(GO_FLAGS) $(MAIN_PACKAGE)

# Run the application
.PHONY: run
run:
	@echo ">> Running application..."
	go run $(MAIN_PACKAGE)

# Clean up build artifacts
.PHONY: clean
clean:
	@echo ">> Cleaning up..."
	rm -rf $(BUILD_DIR)

# Show help message
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make deps    - Install dependencies"
	@echo "  make fmt     - Format source code"
	@echo "  make lint    - Run linter (requires golangci-lint)"
	@echo "  make test    - Run tests"
	@echo "  make build   - Build the application"
	@echo "  make run     - Run the application"
	@echo "  make clean   - Clean up build artifacts"
