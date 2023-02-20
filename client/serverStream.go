package main

import (
	"context"
	"io"
	"log"

	pb "github.com/manan-rastogi/learning-grpc/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming has started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil{
		log.Fatalf("Could not send names: %v", err.Error())
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("Error while Streaming: %v", err.Error())
		}
		log.Printf("Message: %v\n", message)
	}

	log.Printf("Streaming has finished")
}