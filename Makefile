.PHONE: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -cover -v ./...