.PHONY: all
all: kind

.PHONY: clean
clean: kind/clean

.PHONY: kind
kind:
	kind create cluster --config kind.yaml

.PHONY: kind/clean
kind/clean:
	kind delete cluster --name distributed-calculator
