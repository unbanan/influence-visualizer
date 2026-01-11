SERVER_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: build



build: 
	rm -rf ${SERVER_DIR}bin
	mkdir -p ${SERVER_DIR}bin
	cd ${SERVER_DIR}src && go build -o ${SERVER_DIR}bin/server cmd/service/main.go