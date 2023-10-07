build:
	go build -o bin/ ./...

format:
	gofumpt -w -extra .

lint:
	golangci-lint run

tidy:
	go mod tidy

check: format tidy lint

setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest
