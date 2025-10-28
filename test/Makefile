.PHONY: test test-unit test-integration test-config clean install help

# Default target
.DEFAULT_GOAL := help

## help: Display this help message
help:
	@echo "Terraform Stacks Test Suite"
	@echo ""
	@echo "Available targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## install: Download Go dependencies
install:
	go mod download
	go mod tidy

## test: Run all tests with timeout
test:
	go test -v -timeout 30m ./...

## test-unit: Run unit tests only
test-unit:
	go test -v -timeout 15m -run "Test(Networking|Compute|App)Workspace" ./...

## test-integration: Run integration tests only
test-integration:
	go test -v -timeout 30m -run "TestFullStackIntegration|TestWorkspaceDependencyOrder" ./...

## test-config: Run configuration tests (no AWS deployment)
test-config:
	go test -v -timeout 5m -run "TestStackConfig" ./...

## test-parallel: Run tests in parallel (4 workers)
test-parallel:
	go test -v -parallel 4 -timeout 30m ./...

## test-quick: Run a quick smoke test
test-quick:
	go test -v -timeout 10m -run "TestStackConfigDefinesCorrectDependencies" ./...

## clean: Clean up test artifacts and cache
clean:
	go clean -testcache
	rm -rf .terraform.lock.hcl
	find .. -type d -name ".terraform" -exec rm -rf {} + 2>/dev/null || true
	find .. -name "terraform.tfstate*" -delete 2>/dev/null || true

## fmt: Format Go code
fmt:
	go fmt ./...

## vet: Run go vet
vet:
	go vet ./...
