package main

import (
	"log"
	"time"
	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

func (s *HelloServer) SayHelloServerStream(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("got the request Names : %v", req.Names);
	for _, name := range req.Names{
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err;
		}
		time.Sleep(2 *time.Second)
	}
	return nil
}



