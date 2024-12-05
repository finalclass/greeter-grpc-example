.PHONY: all
all: build

.PHONY: proto
proto:
	mkdir -p ./pkg/pb
	protoc --proto_path=./contract \
		--go_out=./pkg/pb --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/pb --go-grpc_opt=paths=source_relative \
		greeter.proto

.PHONY: run
run: proto
	go run main.go

.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 GOAMD64=v2 \
	go build -trimpath -ldflags="-s -w" -o bin/greeter .

.PHONY: clean
clean:
	@go clean
	@rm -f bin/*
