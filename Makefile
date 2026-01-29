.PHONY: build
test:
	go test -v -timeout 30s ./... 

build: test
	go build -v ./cmd/book-helper

run: build
	./book-helper


.DEFAULT_GOAL: build