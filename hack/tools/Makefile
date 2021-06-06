# https://github.com/kubernetes/kubernetes/blob/v1.19.0-beta.0/hack/tools/Makefile

export GO111MODULE=on

GOBIN ?= $(shell go env GOPATH)/bin
GOSRC ?= $(shell go env GOPATH)/src

TOOLS = \
	$(GOBIN)/protoc-gen-grpc-gateway \
	$(GOBIN)/protoc-gen-openapiv2 \
	$(GOBIN)/protoc-gen-go \
	$(GOBIN)/protoc-gen-go-grpc \
	$(GOBIN)/protoc-gen-validate \
	$(GOBIN)/golangci-lint \
	$(GOBIN)/goimports \

SRCS = \
	$(GOSRC)/github.com/googleapis/googleapis \
	$(GOSRC)/github.com/envoyproxy/protoc-gen-validate \

.PHONY: all
all: uninstall clean install

.PHONY: install
install: $(TOOLS) $(SRCS) clean
	@echo "\033[31m"
	@echo "Protocol Buffer Compiler Installation"
	@echo "  For Linux:"
	@echo "    $$ apt install -y protobuf-compiler"
	@echo "  For MacOS:"
	@echo "    $$ brew install protobuf"
	@echo "\033[0m"

.PHONY: uninstall
uninstall:
	rm -f $(TOOLS)

.PHONY: clean
clean:
	rm -rf vendor/

.PHONY: deps
deps:
	go mod vendor

$(GOBIN)/protoc-gen-grpc-gateway: deps
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

$(GOBIN)/protoc-gen-openapiv2: deps
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

$(GOBIN)/protoc-gen-go: deps
	go install google.golang.org/protobuf/cmd/protoc-gen-go

$(GOBIN)/protoc-gen-go: deps
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

$(GOBIN)/protoc-gen-validate: deps
	go install github.com/envoyproxy/protoc-gen-validate

$(GOBIN)/golangci-lint: deps
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

$(GOBIN)/goimports: deps
	go install golang.org/x/tools/cmd/goimports

$(GOSRC)/github.com/googleapis/googleapis: deps
	git clone https://github.com/googleapis/googleapis.git $(GOSRC)/github.com/googleapis/googleapis || true

$(GOSRC)/github.com/envoyproxy/protoc-gen-validate: deps
	go get github.com/envoyproxy/protoc-gen-validate
