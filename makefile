.PHONY: all

all: build

build:
	go build -o ./bin/okctl-mcp-server ./cmd/main.go