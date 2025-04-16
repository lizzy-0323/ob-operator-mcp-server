.PHONY: all

all: build

build:
	go build -o okctl-mcp-server ./cmd/main.go