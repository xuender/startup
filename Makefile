tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cespare/reflex@latest
	go install github.com/jondot/goweight@latest

test:
	go test -race -v ./... -gcflags=all=-l

watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -v ./... -gcflags=all=-l'

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...

lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

build:
	go build -v ./...

coverage:
	go test -v -coverprofile=cover.out -covermode=atomic .
	go tool cover -html=cover.out -o cover.html

weight: tools
	goweight
