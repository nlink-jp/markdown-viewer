MODULE  := github.com/nlink-jp/markdown-viewer
BINARY  := mdv
BIN_DIR := dist

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-s -w -X 'main.version=$(VERSION)'"

# macOS Developer ID signing / notarization (see nlink-jp/.github
# CONVENTIONS.md §Code Signing). Defaults match any Developer ID
# Application cert in the keychain and the org-standard notary
# profile. Builds without these fall back to ad-hoc / un-notarized
# with a one-line warning — see scripts/codesign-darwin.sh.
CODESIGN_IDENTITY ?= Developer ID Application
NOTARY_PROFILE    ?= nlink-jp-notary

# Allow callers to override GOCACHE / GOMODCACHE for environments where the
# default cache directory is not writable (e.g. sandboxes, CI containers).
GOCACHE    ?= $(HOME)/.cache/go-build
GOMODCACHE ?= $(shell go env GOPATH)/pkg/mod

export GOCACHE
export GOMODCACHE

PLATFORMS := \
	darwin/amd64 \
	darwin/arm64 \
	linux/amd64 \
	linux/arm64 \
	windows/amd64

.PHONY: build build-all package test clean help

## build: Build binary for the current OS/Arch → ./dist/mdv
build:
	@mkdir -p $(BIN_DIR)
	go build $(LDFLAGS) -o $(BIN_DIR)/$(BINARY) .
	@scripts/codesign-darwin.sh $(BIN_DIR)/$(BINARY) "$(CODESIGN_IDENTITY)"

## build-all: Cross-compile for all target platforms
build-all:
	@mkdir -p $(BIN_DIR)
	$(foreach platform,$(PLATFORMS),$(call build_platform,$(platform)))

define build_platform
	$(eval OS   := $(word 1,$(subst /, ,$(1))))
	$(eval ARCH := $(word 2,$(subst /, ,$(1))))
	$(eval EXT  := $(if $(filter windows,$(OS)),.exe,))
	$(eval OUT  := $(BIN_DIR)/$(BINARY)-$(OS)-$(ARCH)$(EXT))
	@echo "Building $(OUT)..."
	GOOS=$(OS) GOARCH=$(ARCH) go build $(LDFLAGS) -o $(OUT) .
	@scripts/codesign-darwin.sh $(OUT) "$(CODESIGN_IDENTITY)"

endef

## package: Cross-compile, sign, zip (with LICENSE + README), notarize darwin → dist/
package: build-all
	$(foreach platform,$(PLATFORMS), \
		$(eval OS   := $(word 1,$(subst /, ,$(platform)))) \
		$(eval ARCH := $(word 2,$(subst /, ,$(platform)))) \
		$(eval EXT  := $(if $(filter windows,$(OS)),.exe,)) \
		$(eval BIN  := $(BIN_DIR)/$(BINARY)-$(OS)-$(ARCH)$(EXT)) \
		$(eval ZIP  := $(BIN_DIR)/$(BINARY)-$(VERSION)-$(OS)-$(ARCH).zip) \
		zip -j $(ZIP) $(BIN) LICENSE README.md ;)
	@scripts/notarize-darwin.sh $(BIN_DIR)/$(BINARY)-$(VERSION)-darwin-amd64.zip "$(NOTARY_PROFILE)"
	@scripts/notarize-darwin.sh $(BIN_DIR)/$(BINARY)-$(VERSION)-darwin-arm64.zip "$(NOTARY_PROFILE)"

## test: Run all unit tests
test:
	go test -v -race -count=1 ./...

## clean: Remove build artifacts
clean:
	rm -rf $(BIN_DIR)

## help: Show available targets
help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## //'
