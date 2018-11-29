.PHONY: all lint test

all:
	go build ./...

lint:
	golangci-lint run -E gocyclo -E goimports -E nakedret -E golint --exclude-use-default=false

test:
	go test -v ./... -race