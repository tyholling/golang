ifeq ($(TARGETPLATFORM), linux/amd64)
	GOARCH = amd64
else ifeq ($(TARGETPLATFORM), linux/arm64)
	GOARCH = arm64
endif

build: buf
	CGO_ENABLED=0 go build -o bin/ ./...

buf:
	buf generate

format:
	gofumpt -w -extra .

lint:
	golangci-lint run

tidy:
	go mod tidy

check: format lint tidy

setup:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install mvdan.cc/gofumpt@latest
