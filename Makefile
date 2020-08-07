.PHONY: gela gela-cross evm all test clean
.PHONY: gela-linux gela-linux-386 gela-linux-amd64 gela-linux-mips64 gela-linux-mips64le
.PHONY: gela-darwin gela-darwin-386 gela-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= 1.12
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

gela:
	go run build/ci.go install ./cmd/gela
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gela\" to launch gela."

gc:
	go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	go run build/ci.go install

test: all
	go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

gela-cross: gela-windows-amd64 gela-darwin-amd64 gela-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gela-*

gela-linux: gela-linux-386 gela-linux-amd64 gela-linux-mips64 gela-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gela-linux-*

gela-linux-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gela
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gela-linux-* | grep 386

gela-linux-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gela
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gela-linux-* | grep amd64

gela-linux-mips:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gela
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gela-linux-* | grep mips

gela-linux-mipsle:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gela
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gela-linux-* | grep mipsle

gela-linux-mips64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gela
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gela-linux-* | grep mips64

gela-linux-mips64le:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gela
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gela-linux-* | grep mips64le

gela-darwin: gela-darwin-386 gela-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gela-darwin-*

gela-darwin-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gela
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gela-darwin-* | grep 386

gela-darwin-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gela
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gela-darwin-* | grep amd64

gela-windows-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gela
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gela-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
