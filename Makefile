.PHONY: docker-run
docker-run:
ifdef config
	docker build --build-arg CONFIG=$(config) -q -t yadro-task .
else
	docker build --build-arg CONFIG=test_1.txt -q -t yadro-task .
endif
	docker run --rm -it yadro-task 

.PHONY: build-run
build:
	go build -o ./bin/task ./cmd/task
ifdef config
	./bin/task configs/$(config)
else
	./bin/task
endif

.DEFAULT_GOAL := run

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: count
count:
	find . -name '*.go' | xargs wc -l

.PHONY: count_test
count_test:
	find . -name '*_test.go' | xargs wc -l