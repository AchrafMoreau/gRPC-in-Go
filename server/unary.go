package main

import (
	"context"
	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello World!",
	}, nil
}
