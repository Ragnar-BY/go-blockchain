.PHONY: all lint test

all:
	go build ./cmd/cli

lint:
	golangci-lint run

test:
	go test -v ./... -race