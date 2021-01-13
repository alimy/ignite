PACKGE ?= github.com/alimy/ignite

GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

LDFLAGS += -X "$(PACKGE)/version.BuildTime=$(shell date -v+8H -u '+%Y-%m-%d %H:%M:%S %Z+8')"
LDFLAGS += -X "$(PACKGE)/version.GitHash=$(shell git rev-parse --short=12 HEAD)"

TARGET = ignite
RELEASE_ROOT = release
RELEASE_LINUX_AMD64 = $(RELEASE_ROOT)/linux-amd64/$(TARGET)
RELEASE_DARWIN_AMD64 = $(RELEASE_ROOT)/darwin-amd64/$(TARGET)
RELEASE_WINDOWS_AMD64 = $(RELEASE_ROOT)/windows-amd64/$(TARGET)

.PHONY: build
build: fmt
	go build  -ldflags '$(LDFLAGS)' -o $(TARGET) main.go

.PHONY: install
install: fmt
	go install

.PHONY: release
release: darwin-amd64
	cp LICENSE README.md $(RELEASE_DARWIN_AMD64)
	cd $(RELEASE_DARWIN_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-darwin_amd64.zip $(TARGET) && cd -

.PHONY: linux-amd64
linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags '$(LDFLAGS)' -o $(RELEASE_LINUX_AMD64)/$(TARGET) main.go

.PHONY: darwin-amd64
darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -ldflags '$(LDFLAGS)' -o $(RELEASE_DARWIN_AMD64)/$(TARGET) main.go

.PHONY: windows-x64
windows-x64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -ldflags '$(LDFLAGS)' -o $(RELEASE_WINDOWS_AMD64)/$(TARGET) main.go

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)
