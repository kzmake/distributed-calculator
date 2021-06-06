.PHONY: all
all: pre fmt lint proto

.PHONY: pre
pre:
	go mod tidy

.PHONY: fmt
fmt:
	goimports -l -local "github.com/kzmake/distributed-calculator/common/health" -w .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: proto
proto:
	@for file in api/**/**/*.proto; do \
		protoc \
			--proto_path=.:. \
			--proto_path=../../third_party/proto:. \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			--grpc-gateway_out=logtostderr=true,paths=source_relative:. \
			--openapiv2_out=logtostderr=true:. \
			$$file; \
		echo "generated from $$file"; \
	done
