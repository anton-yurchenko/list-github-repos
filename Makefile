# global
BINARY := $(notdir $(CURDIR))
GO_BIN_DIR := $(GOPATH)/bin
OSES := linux darwin windows
ARCHS := amd64

# initialize
.PHONY: init
init:
	@rm -f go.mod
	@rm -f go.sum
	@rm -rf ./vendor
	@go mod init $$(pwd | awk -F'/' '{print "github.com/"$$(NF-1)"/"$$NF}')

.PHONY: release
release:
	@rm -rf ./release
	@mkdir -p release
	@for ARCH in $(ARCHS); do \
		for OS in $(OSES); do \
			if test "$$OS" = "windows"; then \
				GOOS=$$OS GOARCH=$$ARCH go build -o release/$(BINARY)-$$OS-$$ARCH.exe; \
			else \
				GOOS=$$OS GOARCH=$$ARCH go build -o release/$(BINARY)-$$OS-$$ARCH; \
			fi; \
		done; \
	done
