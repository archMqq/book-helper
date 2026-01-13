.PHONY: build
build:
	go build -0 librarian ./

.DEFAULT_GOAL: build