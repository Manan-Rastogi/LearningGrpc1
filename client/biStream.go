package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/manan-rastogi/learning-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList){
	log.Println("Bidirectional streaming started")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil{
		log.Fatalf("Could not send names: %v", err.Error())
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Fatalf("Error while streaming: %v", err.Error())
			}

			log.Print(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil{
			log.Fatalf("Error while sending: %v", err.Error())
		}
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
	<- waitc
	log.Printf("Bidirectional streaming finished")

}