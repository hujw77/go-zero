# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: kratos kratos-protoc kratos-gen-bts kratos-gen-mc kratos-gen-project   
.PHONY: protoc-gen-bm protoc-gen-bswagger protoc-gen-ecode protoc-gen-gofast
.PHONY: testgen testcli
.PHONY: admin-bbq-comment

# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash
ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"
# We don't need make's built-in rules.
MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

.EXPORT_ALL_VARIABLES:
GOBIN  := $(go env GOBIN)
ifeq ($(GOBIN),)
   GOBIN := ~/go/bin
endif
GO ?= latest
GORUN := env GO111MODULE=on go run
BIN_DIR := ./build/bin

kratos:
	$(GORUN) build/ci.go install ./tool/kratos
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/kratos\" to launch."

kratos-protoc:
	$(GORUN) build/ci.go install ./tool/kratos-protoc
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/kratos-protoc\" to launch."

kratos-gen-bts:
	$(GORUN) build/ci.go install ./tool/kratos-gen-bts
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/kratos-gen-bts\" to launch."

kratos-gen-mc:
	$(GORUN) build/ci.go install ./tool/kratos-gen-mc
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/kratos-gen-mc\" to launch."

kratos-gen-project:
	$(GORUN) build/ci.go install ./tool/kratos-gen-project
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/kratos-gen-project\" to launch."

protoc-gen-bm:
	$(GORUN) build/ci.go install ./tool/protobuf/protoc-gen-bm
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/protoc-gen-bm\" to launch."

protoc-gen-bswagger:
	$(GORUN) build/ci.go install ./tool/protobuf/protoc-gen-bswagger
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/protoc-gen-bswagger\" to launch."

protoc-gen-ecode:
	$(GORUN) build/ci.go install ./tool/protobuf/protoc-gen-ecode
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/protoc-gen-ecode\" to launch."

protoc-gen-gofast:
	env GO111MODULE=on GOBIN= go get -u github.com/gogo/protobuf/protoc-gen-gofast
	@echo "Run \"$(GOBIN)/protoc-gen-gofast\" to launch."

testgen:
	$(GORUN) build/ci.go install ./tool/testgen
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/testgen\" to launch."

testcli:
	$(GORUN) build/ci.go install ./tool/testcli
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/testcli\" to launch."

devtools: kratos kratos-protoc kratos-gen-bts kratos-gen-mc kratos-gen-project \
	protoc-gen-bm protoc-gen-bswagger protoc-gen-ecode protoc-gen-gofast \
	testgen testcli
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

install:
	@cp -vi $(BIN_DIR)/* $(GOBIN) 


admin-bbq-comment:
	$(GORUN) build/ci.go install ./admin/bbq/comment/cmd
	@echo "Done building."
	@echo "Run \"$(BIN_DIR)/admin-bbq-comment\" to launch."

