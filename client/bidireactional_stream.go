package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/akshtrikha/grpc-demo-golang/proto"
)

func callHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send the names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Printf("Error while streaming: %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the request: %v", err)
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc

	log.Printf("Bidirectional Streaming")
}
