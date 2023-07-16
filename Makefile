.PHONE: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -cover -count=5 ./...

.PHONY: test-once
test-once:
	go test -cover -count=1 ./...

