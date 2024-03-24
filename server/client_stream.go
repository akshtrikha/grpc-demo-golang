package main

import (
	"io"
	"log"

	pb "github.com/akshtrikha/grpc-demo-golang/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	log.Printf("Got reqest to SayHelloClientStreaming")

	var names []string

	for {
		payload, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: names})
			// break
		} else if err != nil {
			log.Fatalf("Error while getting the stream from the client: %v", err)
			return err
		}

		log.Printf("payload in the stream: %v", payload.Name)
		names = append(names, payload.Name)
	}
}
