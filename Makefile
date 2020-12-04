GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
GOFMT=gofmt
GOGEN=go generate
GOINSTALL=go install
GOIMPORTS=goimports
GOLINT=golangci-lint
BINARY_NAME=todo
CMD_PKG=./cmd/todo
SMOKE_DIR=./test/smoke

all: help

.PHONY: init
init: ## initilize developper environment
	go install github.com/golang/mock/mockgen

.PHONY: get
get: ## go get dependencies
	$(GOGET) -v -t ./...

.PHONY: install
install: ## go install
	$(GOINSTALL) -v $(CMD_PKG)

.PHONY: build
build: ## build go binary
	$(GOBUILD) -o $(BINARY_NAME) -v $(CMD_PKG)

.PHONY: mockgen
mockgen: ## generate mock
	$(GOGEN) ./...

.PHONY: test
test: ## go test
	$(GOTEST) -v ./...

.PHONY: smoke
smoke: ## run smoke test
	$(GOTEST) -v $(SMOKE_DIR) -tags smoke

.PHONY: clean
clean: ## remove go binary
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: fmt
fmt: ## format go files
	$(GOFMT) -l -w -s .
	$(GOIMPORTS) -w .

.PHONY: lint
# need docker to run this command
# this command just run golangci-lint
# so, if you hate docker, you can run equivalent this installing golangci-lint locally
lint: ## check lint, format
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.31.0 golangci-lint run -v

uninstall: ## uninstall todo-cli
	./scripts/uninstall.sh

.PHONY: help
help: ## DIsplay this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
