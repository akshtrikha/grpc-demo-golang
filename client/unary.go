package main

import (
	"context"
	"log"
	"time"

	pb "github.com/akshtrikha/grpc-demo-golang/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Println("Calling client.SayHello")
	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Logging the response from the server!")
	log.Printf("%s", res.Message)
}
