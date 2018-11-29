.PHONY: all lint

all:
	go build ./...

lint:
	golangci-lint run -E gocyclo -E goimports -E nakedret
