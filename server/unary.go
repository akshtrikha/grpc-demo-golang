package main

import (
	"context"

	pb "github.com/akshtrikha/grpc-demo-golang/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from the server! This is a unary connection!!",
	}, nil
}
