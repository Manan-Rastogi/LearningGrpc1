package main

import (
	"context"
	"log"
	"time"

	pb "github.com/manan-rastogi/learning-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil{
		log.Fatalf("Could not greet: %v", err.Error())
	}

	log.Printf("Response: %v", res.Message)
}