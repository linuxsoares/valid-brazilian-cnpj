export GOBIN := $(PWD)/tools
GOLANGCI_VERSION := v1.60.1
export GOLANGCI_TIMEOUT := 8m
golangci-lint := $(GOBIN)/golangci-lint-$(GOLANGCI_VERSION)
export GOLANGCI_LINT_BIN := $(golangci-lint)
$(golangci-lint):
	@echo "Installing golangci-lint $(GOLANGCI_VERSION)..." && \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b tools $(GOLANGCI_VERSION) && \
	mv $(GOBIN)/golangci-lint $(GOBIN)/golangci-lint-$(GOLANGCI_VERSION)

test:
	go test ./...

lint: $(golangci-lint)
	$(GOBIN)/golangci-lint-$(GOLANGCI_VERSION) run
