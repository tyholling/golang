build: tidy
	go build -o bin/ ./...

format:
	gofumpt -w -extra .

lint:
	golangci-lint run

tidy:
	go mod tidy

vuln:
	govulncheck ./...

check: format lint

setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install mvdan.cc/gofumpt@latest
