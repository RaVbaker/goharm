GOCMD:=go
BUILD_DIR:=build
GOBUILD:=$(GOCMD) build
GOCLEAN:=$(GOCMD) clean
GOTEST:=$(GOCMD) test
GOARCH:=amd64
PLATFORMS:=linux darwin
GOOS=$(word 1, $@)
BINARY_NAME=goharm
LDFLAGS=-ldflags "-X 'main.Version=v$(VERSION)'"
VERSION?=$(shell git describe --tags --always --dirty)

all: test build

$(PLATFORMS):
	mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH) -v ./cmd/goharm

build: linux darwin

test:
    # 	go get -u github.com/rakyll/gotest
	$(GOTEST) -race -timeout 100s -cover -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

tag:
	if test "$(TAG)" = "" ; then \
		echo "usage: make tag TAG=1.2.3"; \
		exit 1; \
	fi
	git tag -a $(TAG) -m "$(TAG)"
	git push origin $(TAG)

install:
	$(GOCMD) install github.com/ravbaker/goharm/cmd/...

lint:
	go vet ./...

.PHONY: all build test clean tag $(PLATFORMS)