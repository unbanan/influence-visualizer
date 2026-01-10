PROTO_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
GEN_GO_DIR := $(PROTO_DIR)go
GEN_CPP_DIR := $(PROTO_DIR)cpp
GRPC_CPP_PLUGIN := $(shell which grpc_cpp_plugin)

# Add GOPATH bin to PATH for protoc plugins
export PATH := $(shell go env GOPATH)/bin:$(PATH)

.PHONY: install go cpp all

install:
	@echo "Установка protoc плагинов..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

go:
	rm -rf $(GEN_GO_DIR)
	mkdir -p ${GEN_GO_DIR}
	protoc -I=${PROTO_DIR} \
	--go_out=${GEN_GO_DIR} \
	--go-grpc_out=${GEN_GO_DIR} \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	*.proto

cpp:
	rm -rf $(GEN_CPP_DIR)
	mkdir -p ${GEN_CPP_DIR}
	protoc -I=${PROTO_DIR} \
	--cpp_out=$(GEN_CPP_DIR) \
	--grpc_out=$(GEN_CPP_DIR) \
	--plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN) \
	*.proto

all: go cpp
