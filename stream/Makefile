grpc:
	@protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto

build-server:
	@go build -v -o server server/server.go

run-server: build-server
	@./server/server
