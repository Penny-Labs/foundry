# Variables
COVERAGE_FILE=coverage.out

.PHONY: test
test:
	@go test ./... -coverprofile=$(COVERAGE_FILE)

.PHONY: testv
testv:
	@go test ./... -v -coverprofile=$(COVERAGE_FILE)

.PHONY: cover
cover:
	@go tool cover -html=$(COVERAGE_FILE)
    @echo "Test coverage report generated: $(COVERAGE_FILE) and coverage.html"

.PHONY: lint
lint:
	@golangci-lint run ./...