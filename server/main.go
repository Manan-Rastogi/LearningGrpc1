package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/manan-rastogi/learning-grpc/proto"
)

const (
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("Failed to start server: %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to start: %v", err.Error())
	}
	
}