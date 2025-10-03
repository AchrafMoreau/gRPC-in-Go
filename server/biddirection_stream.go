package main

import (
	"io"
	"log"
	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

func (s *HelloServer) SayHelloBiDirectionStreaming(stream pb.GreetService_SayHelloBiDirectionStreamingServer) error{

	for{
		req, err := stream.Recv();
		if err == io.EOF{
			return nil;
		}
		if err != nil{
			log.Fatalf("Error Receving Stream Names :%v", err)
			return err
		}

		log.Printf("Reciving Name : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil{
			log.Fatalf("Error Streaming Names Back :%v", err)
			return err;
		}
		log.Printf("Sending Name : %v", req.Name)
	}

}
