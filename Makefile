GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

LDFLAGS += -X "github.com/alimy/ignite/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/ignite/version.GitHash=$(shell git rev-parse HEAD)"

TARGET = ignite
RELEASE_ROOT = release
RELEASE_LINUX_AMD64 = $(RELEASE_ROOT)/linux-amd64/$(TARGET)
RELEASE_DARWIN_AMD64 = $(RELEASE_ROOT)/darwin-amd64/$(TARGET)
RELEASE_WINDOWS_AMD64 = $(RELEASE_ROOT)/windows-amd64/$(TARGET)

.PHONY: build
build: fmt
	go build  -ldflags '$(LDFLAGS)' -o $(TARGET) main.go

.PHONY: release
release: linux-amd64 darwin-amd64 windows-x64
	cp LICENSE README.md $(RELEASE_LINUX_AMD64)
	cp LICENSE README.md $(RELEASE_DARWIN_AMD64)
	cp LICENSE README.md $(RELEASE_WINDOWS_AMD64)
	cd $(RELEASE_LINUX_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-linux_amd64.zip $(TARGET) && cd -
	cd $(RELEASE_DARWIN_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-darwin_amd64.zip $(TARGET) && cd -
	cd $(RELEASE_WINDOWS_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-windows_amd64.zip $(TARGET) && cd -

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
