# Version is derived from tags
VERSION ?= $(shell git describe --dirty --always --tags)

# REGISTRY_HOST defines the registry and organization used to publish the images.
# Use localhost:5000 for the local registry
REGISTRY_HOST ?= docker.io/koorinc

# Image URL to use for all building/pushing image targets
IMG ?= $(REGISTRY_HOST)/version-service:$(VERSION)

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: generate
generate: buf
	$(BUF) build
	$(BUF) generate

.PHONY: lint
lint: buf
	$(BUF) lint
	$(BUF) breaking --against '.git#branch=main'

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: local-certs
local-certs: certs/cert.pem ## generate local certs

certs/cert.pem:
	mkcert -install -cert-file=certs/cert.pem -key-file=certs/key.pem localhost

##@ Build

.PHONY: build
build: generate fmt vet ## Build server binary.
	go build -o bin/server main.go

.PHONY: run
run: generate fmt vet ## Run a server from your host.
	@echo "Listening on http://localhost:8082"
	NO_TLS=true go run ./main.go

.PHONY: run-tls
run-tls: generate fmt vet local-certs ## Run a server from your host.
	@echo "Listening on https://localhost:8082"
	go run ./main.go

# If you wish to build the version service image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64 ). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
.PHONY: docker-build
docker-build: ## Build docker image with the version service.
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the version service.
	docker push ${IMG}

.PHONY: docker-run-it
docker-run-it: docker-build ## Run docker image
	docker run -it --rm -p 8082:8082 -e NO_TLS=true ${IMG}

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
BUF ?= $(LOCALBIN)/buf

## Tool Versions
BUF_VERSION ?= v1.26.1

.PHONY: buf
buf: $(BUF)
$(BUF):
	GO111MODULE=on GOBIN=$(LOCALBIN) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)