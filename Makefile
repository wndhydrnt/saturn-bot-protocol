ARCH=$(shell uname -m)
BINARY=integration-test
OS=$(shell uname -s)
VERSION?=0.0.0-dev

build:
	go test -c -o $(BINARY)

build_all: build_darwin_amd64 build_darwin_arm64 build_linux_arm64 build_linux_amd64 checksums

build_darwin_amd64:
	GOARCH=amd64 GOOS=darwin go test -c -o $(BINARY)-$(VERSION).Darwin-amd64

build_darwin_arm64:
	GOARCH=arm64 GOOS=darwin go test -c -o $(BINARY)-$(VERSION).Darwin-arm64

build_linux_arm64:
	GOARCH=arm64 GOOS=linux go test -c -o $(BINARY)-$(VERSION).Linux-aarch64

build_linux_amd64:
	GOARCH=amd64 GOOS=linux go test -c -o $(BINARY)-$(VERSION).Linux-x86_64

checksums:
	sha256sum $(BINARY)-$(VERSION).* > sha256sums.txt
