build: buf tidy
	go build -o bin/ ./...

buf:
	buf lint
	buf generate

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
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install mvdan.cc/gofumpt@latest
