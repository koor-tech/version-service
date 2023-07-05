.PHONY: build
build: buf
	$(BUF) generate

lint: buf
	$(BUF) lint
	$(BUF) breaking --against '.git#branch=main'

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
BUF ?= $(LOCALBIN)/buf

## Tool Versions
BUF_VERSION ?= v1.23.1

.PHONY: buf
buf: $(BUF)
$(BUF):
	GO111MODULE=on GOBIN=$(LOCALBIN) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)