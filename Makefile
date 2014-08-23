.PHONY: default all clean build run deploy save_deps
default: all

all: build

build: clean
	gox -osarch="windows/amd64 linux/amd64 darwin/amd64" -output="build/rcmd-{{.OS}}-{{.Arch}}"

clean:
	rm -rf  build

run: build
	build/rcmd-darwin-amd64 --config config.yml

save_deps:
	godep save

.PHONY: dev_bootstrap
dev_bootstrap:
	go get ./...
	go get github.com/mitchellh/gox
	gox -build-toolchain -osarch="darwin/amd64" -osarch="linux/amd64" -osarch="windows/amd64"

