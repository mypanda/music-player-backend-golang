.PHONY: build clean tool lint help test

PROJECTNAME=$(shell basename "$(PWD)")

all: build

run:
	go run .

build:
	go build -v .

tool:
	go tool vet . |&amp; grep -v vendor; true
	gofmt -w .

lint:
	golint ./...

test:
	go test 

clean:
	rm -rf ${PROJECTNAME}
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v