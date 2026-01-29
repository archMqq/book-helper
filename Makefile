.PHONY: build
test:
	go test -v -timeout 30s ./... 

build: test
	go build -v -o gptm ./cmd/gpt

run: build
	./gptm


.DEFAULT_GOAL: build