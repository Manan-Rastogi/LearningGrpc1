package main

import (
	"context"
	"log"
	"time"

	pb "github.com/manan-rastogi/learning-grpc/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList){
	log.Println("Client Streaming Started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil{
		log.Fatalf("could not send names: %v", err.Error())
	}

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil{
			log.Fatalf("Error while sending: %v", err.Error())
		}
		log.Printf("Send Request with name: %s", name)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("Error while receiving: %v", err.Error())
	}
	log.Println("Client Streaming Finished")

	log.Printf("%v", res.Message)
} 