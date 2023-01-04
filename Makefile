
build:
	go build -v ./...

test:
	go test ./... -race

fmt:
	go fmt ./...

.PHONY: build test fmt
