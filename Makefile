.PHONY: run test test-watch build push

# The binary to build (just the basename).
BIN := git-ops-update

# Where to push the docker image.
REGISTRY ?= ghcr.io/choffmeister

IMAGE := $(REGISTRY)/$(BIN)
VERSION := test

MAIN := ./cmd/git-ops-update
TEST := ./internal

run:
	go run $(MAIN) --dir=example --dry

build:
	mkdir -p build/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/git-ops-update-linux-amd64 $(MAIN)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/git-ops-update-darwin-amd64 $(MAIN)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/git-ops-update-windows-amd64.exe $(MAIN)

test:
	go test -v $(TEST)

test-watch:
	watch -n1 go test -v $(TEST)

test-cover:
	go test -coverprofile=coverage.out $(TEST)
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

container:
	docker build -t $(IMAGE):$(VERSION) .

container-push: container
	docker push $(IMAGE):$(VERSION)
