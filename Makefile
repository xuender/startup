tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cespare/reflex@latest

test:
	go test -race -v ./... -gcflags=all=-l

watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -v ./...'

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...

lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...
