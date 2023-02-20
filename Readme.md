# Readme for grpc

## Installing Protoc
```
https://grpc.io/docs/languages/go/quickstart/
```

## Install on Windows
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Run a protoc file
```
protoc --go_out=. --go-grpc_out=. proto/greet.proto

                        OR

make create_grpc
```
