.DEFAULT_GOAL: help

.PHONY: help
help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run HTTP server locally on port 8080
	@go run ./cmd/pod-overlap-service

.PHONY: build
build: ## Build the Go binary for production
	@mkdir -p build
	@go build -ldflags="-s -w" -o build/pod-overlap-service ./cmd/pod-overlap-service

.PHONY: test
test: ## Execute the tests in the development environment
	@go test ./... -count=1 -timeout 2m

.PHONY: coverage
coverage: ## Generate test coverage in the development environment
	go test ./... -coverprofile=/tmp/coverage.out -coverpkg=./...
	go tool cover -html=/tmp/coverage.out

.PHONY: lint
lint: ## Execute syntactic analysis and autofix minor problems
	@golangci-lint run --fix
