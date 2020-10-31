GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
GOINSTALL=go install
GOIMPORTS=goimports
GOLINT=golangci-lint
BINARY_NAME=todo
CMD_PKG=./cmd/todo
SMOKE_DIR=./test/smoke

all: help

get: ## go get dependencies
	$(GOGET) -v -t ./...

install: ## go install
	$(GOINSTALL) -v $(CMD_PKG)

build: ## build go binary
	$(GOBUILD) -o $(BINARY_NAME) -v $(CMD_PKG)

.PHONY: test
test: ## go test
	$(GOTEST) -v ./...

.PHONY: smoke
smoke: ## run smoke test
	$(GOTEST) -v $(SMOKE_DIR) -tags smoke

.PHONY: clean
clean: ## remove go binary
	$(GOCLEAN)
	rm $(BINARY_NAME)

fmt: ## format go files
	$(GOFMT) -l -w -s .
	$(GOIMPORTS) -w .

.PHONY: lint
lint: ## check lint, format
	$(GOLINT) run

.PHONY: help
help: ## DIsplay this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
