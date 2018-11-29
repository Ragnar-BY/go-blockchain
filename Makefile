.PHONY: all lint test

all:
	go build ./...

lint:
	golangci-lint run

test:
	go test -v ./... -race