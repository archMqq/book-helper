.PHONY: build
build:
	go build -v ./cmd/book-helper

.DEFAULT_GOAL: build