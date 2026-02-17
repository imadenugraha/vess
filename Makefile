SHELL := /bin/sh

BINARY ?= vess
GO ?= go
PKG ?= ./...

DOCKERFILE ?= Dockerfile
IMAGE_TAG ?= $(BINARY):latest
OS ?= alpine
PHP_VERSION ?= 8.3
TYPE ?= fpm
ENV_FILE ?= examples/basic.env

.PHONY: help build run test test-race fmt vet tidy clean install \
	generate docker-build dev-check

help: ## Show available targets
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_.-]+:.*##/ {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build CLI binary (default: vess)
	$(GO) build -o $(BINARY) .

run: ## Run the CLI (use ARGS="...")
	$(GO) run . $(ARGS)

test: ## Run all tests
	$(GO) test $(PKG)

test-race: ## Run tests with race detector
	$(GO) test -race $(PKG)

fmt: ## Format Go code
	$(GO) fmt $(PKG)

vet: ## Run go vet checks
	$(GO) vet $(PKG)

tidy: ## Sync go.mod/go.sum
	$(GO) mod tidy

clean: ## Remove built binary
	rm -f $(BINARY)

install: ## Install CLI to GOPATH/bin
	$(GO) install .

generate: build ## Generate Dockerfile from env config
	./$(BINARY) generate -o $(OS) -p $(PHP_VERSION) -t $(TYPE) -e $(ENV_FILE) -f $(DOCKERFILE)

docker-build: ## Build Docker image from Dockerfile
	docker build -f $(DOCKERFILE) -t $(IMAGE_TAG) .

dev-check: fmt vet test ## Run local CI checks (fmt, vet, test)
