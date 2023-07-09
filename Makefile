.PHONY: build
build:
	go build -v ./cmd/task

.DEFAULT_GOAL := build

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: count
count:
	find . -name '*.go' | xargs wc -l

.PHONY: count_test
count_test:
	find . -name '*_test.go' | xargs wc -l