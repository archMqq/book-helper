.PHONY: build
build:
	go build -v ./cmd/book-helper

run: build
	./book-helper
.DEFAULT_GOAL: build