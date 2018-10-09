TEST ?= $(shell go list ./... | grep -v -e vendor -e keys -e tmp)
BUILD=tmp/bin

INFO_COLOR=\033[1;34m
RESET=\033[0m
BOLD=\033[1m

default: build

depsdev: ## Installing dependencies for development
	go get github.com/golang/lint/golint

test: ## Run test
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Testing$(RESET)"
	go test -v $(TEST) -timeout=30s -parallel=4
	go test -race $(TEST)

lint: ## Exec golint
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Linting$(RESET)"
	golint -min_confidence 1.1 -set_exit_status $(TEST)

build:
	go build -o $(BUILD)/patt
