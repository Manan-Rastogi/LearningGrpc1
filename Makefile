create_grpc: 
	protoc --go_out=. --go-grpc_out=. proto/greet.proto

.PHONY: create_grpc