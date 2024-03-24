package main

import (
	"context"
	"io"
	"log"

	pb "github.com/akshtrikha/grpc-demo-golang/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Streaming Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Could not send names; Error: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error while streaming. Error:	 %v", err)
		}

		log.Println(message)
	}

	log.Printf("Streaming finished.")
}

