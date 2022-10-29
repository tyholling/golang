build:
	go build -o bin/ ./...

align:
	find . -name *.go ! -name *.pb.go -exec fieldalignment {} \;

buf:
	buf generate

critic:
	gocritic check -enableAll ./...

format:
	gofumpt -w -extra .

revive:
	revive -formatter stylish ./...

secure:
	gosec -quiet -show-ignored ./...

staticcheck:
	staticcheck -show-ignored ./...

tidy:
	go mod tidy

vet:
	go vet ./...

check: align buf critic format revive secure staticcheck tidy vet

setup:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/go-critic/go-critic/cmd/gocritic@latest
	go install github.com/mgechev/revive@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install mvdan.cc/gofumpt@latest

images:
	docker build -t client -f deploy/client .
	docker build -t server -f deploy/server .
