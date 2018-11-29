.PHONY: all lint test

all:
	go build ./...

lint:
	golangci-lint run -E gocyclo -E goimports -E nakedret

test:
	go test -v ./... -race