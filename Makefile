SERVICES := \
	microservices/adder \
	microservices/subtractor \

.PHONY: all
all: pre fmt lint proto

.PHONY: pre
pre:
	@for f in $(SERVICES); do make -C $$f pre; done

.PHONY: fmt
fmt:
	@for f in $(SERVICES); do make -C $$f fmt; done

.PHONY: lint
lint:
	@for f in $(SERVICES); do make -C $$f lint; done

.PHONY: proto
proto:
	@for file in api/**/**/*.proto; do \
		protoc \
			--proto_path=.:. \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			--grpc-gateway_out=logtostderr=true,paths=source_relative:. \
			--openapiv2_out=logtostderr=true:. \
			$$file; \
		echo "generated from $$file"; \
	done

.PHONY: kind
kind:
	kind get clusters -q | grep "distributed-calculator" || kind create cluster --config kind.yaml

.PHONY: clean
clean:
	kind delete cluster --name distributed-calculator

.PHONY: dev
dev:
	skaffold dev

.PHONY: http
http:
	curl -i localhost:58080/add -H "Content-Type: application/json" -d '{"operand_one": "10", "operand_two": "0.123"}'
	curl -i localhost:58080/sub -H "Content-Type: application/json" -d '{"operand_one": "10", "operand_two": "0.123"}'

.PHONY: grpc
grpc:
	grpcurl -plaintext -d '{"operand_one": "10", "operand_two": "0.123"}' -proto api/adder/v1/adder.proto localhost:50051 calculator.adder.v1.Adder/Add
	grpcurl -plaintext -d '{"operand_one": "10", "operand_two": "0.123"}' -proto api/subtractor/v1/subtractor.proto localhost:50051 calculator.subtractor.v1.Subtractor/Sub
