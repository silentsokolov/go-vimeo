# Check style and linting
.PHONY: lint

lint:
	@echo "--> Running golangci"
	@golangci-lint run ./vimeo

# Format code
.PHONY: fmt

fmt:
	@echo "--> Running go fmt"
	@go fmt ./vimeo

# Test
.PHONY: test

test:
	@echo "--> Running tests"
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./vimeo
