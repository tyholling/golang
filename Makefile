build: buf
	go build -o bin/ ./...

align:
	find . -name *.go ! -name *.pb.go -exec fieldalignment {} \;

ascii:
	asciicheck ./...

buf:
	buf generate

critic:
	gocritic check -enableAll ./...

errcheck:
	errcheck -ignoregenerated ./...

format:
	gofumpt -w -extra .

nilness:
	nilness ./...

revive:
	revive -formatter stylish ./...

secure:
	gosec -quiet ./...

shadow:
	shadow ./...

staticcheck:
	staticcheck -show-ignored ./...

tidy:
	go mod tidy

vet:
	go vet ./...

vuln:
	govulncheck ./...

check: align ascii critic errcheck format nilness revive secure shadow staticcheck tidy vet

setup:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/go-critic/go-critic/cmd/gocritic@latest
	go install github.com/kisielk/errcheck@latest
	go install github.com/mgechev/revive@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/tdakkota/asciicheck/cmd/asciicheck@latest
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	go install golang.org/x/tools/go/analysis/passes/nilness/cmd/nilness@latest
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
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
	podman build -t builder -f docker/builder .
