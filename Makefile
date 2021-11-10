.PHONY: run test test-watch build push

# The binary to build (just the basename).
BIN := git-ops-update

# Where to push the docker image.
REGISTRY ?= choffmeister

IMAGE := $(REGISTRY)/$(BIN)

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)

run:
	go run ./src

build:
	mkdir -p build/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/git-ops-update-linux-amd64 ./src
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/git-ops-update-darwin-amd64 ./src
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/git-ops-update-windows-amd64.exe ./src

test:
	go test -v ./src

test-watch:
	watch -n1 go test -v ./src

test-cover:
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

container:
	docker build -t $(IMAGE):$(VERSION) .

container-push: container
	docker push $(IMAGE):$(VERSION)
	docker tag $(IMAGE):$(VERSION) $(IMAGE):latest
	docker push $(IMAGE):latest
