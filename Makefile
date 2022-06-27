ENV_DEV_FILE := .env
ENV_DEV = $(shell cat $(ENV_DEV_FILE))
ENV_TEST_FILE := .env
ENV_TEST = $(shell cat $(ENV_TEST_FILE))

# Lint, Format
.PHONY: lint
lint: src/tools
	golangci-lint run ./...

.PHONY: format
format: src/tools
	golangci-lint run ./... --fix

.PHONY: test
test:
	$(ENV_TEST) go test -v ./...

.PHONY: test-coverage
test-coverage:
	$(ENV_TEST) go test -v ./... -covermode=count

.PHONY: check
check:
	echo "called"
