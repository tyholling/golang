build: buf
	go build -o bin/ ./...

buf:
	buf generate

format:
	gofumpt -w -extra .

lint:
	golangci-lint run

tidy:
	go mod tidy

vuln:
	govulncheck ./...

check: format lint tidy

setup:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install mvdan.cc/gofumpt@latest

images:
	podman build -t localhost:5000/client -f docker/client .
	podman build -t localhost:5000/server -f docker/server .
	podman build -t localhost:5000/http -f docker/http .

push:
	podman push localhost:5000/client
	podman push localhost:5000/server
	podman push localhost:5000/http

builder:
	podman build -t localhost:5000/builder -f docker/builder .
	podman push localhost:5000/builder
