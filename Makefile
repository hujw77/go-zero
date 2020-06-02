# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: kratos kratos-protoc kratos-gen-bts kratos-gen-mc kratos-gen-project   
.PHONY: protoc-gen-bm protoc-gen-bswagger protoc-gen-ecode protoc-gen-gofast
.PHONY: testgen testcli
.PHONY: admin-bbq-comment

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

kratos:
	$(GORUN) build/ci.go install ./tool/kratos
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kratos\" to launch."

kratos-protoc:
	$(GORUN) build/ci.go install ./tool/kratos-protoc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kratos-protoc\" to launch."

kratos-gen-bts:
	$(GORUN) build/ci.go install ./tool/kratos-gen-bts
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kratos-gen-bts\" to launch."

kratos-gen-mc:
	$(GORUN) build/ci.go install ./tool/kratos-gen-mc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kratos-gen-mc\" to launch."

kratos-gen-project:
	$(GORUN) build/ci.go install ./tool/kratos-gen-project
	@echo "Done building."
	@echo "Run \"$(GOBIN)/kratos-gen-project\" to launch."

protoc-gen-bm:
	$(GORUN) build/ci.go install ./tool/protobuf/protoc-gen-bm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/protoc-gen-bm\" to launch."

protoc-gen-bswagger:
	$(GORUN) build/ci.go install ./tool/protobuf/protoc-gen-bswagger
	@echo "Done building."
	@echo "Run \"$(GOBIN)/protoc-gen-bswagger\" to launch."

protoc-gen-ecode:
	$(GORUN) build/ci.go install ./tool/protobuf/protoc-gen-ecode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/protoc-gen-ecode\" to launch."

protoc-gen-gofast:
	env GOBIN= go get -u github.com/gogo/protobuf/protoc-gen-gofast
	@echo "Run \"protoc-gen-gofast\" to launch."

testgen:
	$(GORUN) build/ci.go install ./tool/kratos/testgen
	@echo "Done building."
	@echo "Run \"$(GOBIN)/testgen\" to launch."

testcli:
	$(GORUN) build/ci.go install ./tool/kratos/testcli
	@echo "Done building."
	@echo "Run \"$(GOBIN)/testcli\" to launch."

admin-bbq-comment:
	$(GORUN) build/ci.go install ./admin/bbq/comment/cmd
	@echo "Done building."
	@echo "Run \"$(GOBIN)/admin-bbq-comment\" to launch."

