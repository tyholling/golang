build:
	go build -o bin/ ./...

format:
	gofumpt -w -extra .

setup:
	go install mvdan.cc/gofumpt@latest
