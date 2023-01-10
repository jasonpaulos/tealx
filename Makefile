
build:
	go build -o bin/tealx -v ./language/cli

test:
	go test ./... -race

fmt:
	go fmt ./...

.PHONY: build test fmt
