default: test

test:
	go test ./...

verbose:
	go test -v ./...

fmt:
	find . -type f -name \*.go | xargs gofmt -w

.PHONY: test verbose fmt
