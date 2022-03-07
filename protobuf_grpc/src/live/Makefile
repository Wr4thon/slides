.PHONY: all clean proto mod build docker

all: clean proto mod build docker

clean:
	rm -rf todo_grpc.pb.go todo.pb.go 2>/dev/null || true
	go clean

generate: proto mod

proto:
	protoc -I . todo.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.

mod:
	ls go.mod 2>/dev/null || go mod init
	go mod tidy
	go mod download

build-client: clean proto mod
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
	go build -installsuffix cgo -o todoClient ./cmd/todoClient

build-service: clean proto mod
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
	go build -installsuffix cgo -o service ./cmd/service

docker: COMMIT=${shell git log -1 --pretty=format:"%H"}

docker: clean proto mod build-service
	docker build --build-arg=COMMIT=${COMMIT} -t wr4thon/todo:${COMMIT} -t wr4thon/todo:latest .

structure:
	mkdir -p cmd/service cmd/todoClient
	touch cmd/service/main.go
	touch cmd/todoClient/main.go
	touch todo.proto
