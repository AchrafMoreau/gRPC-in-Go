package main

import (
	"io"
	"log"
	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

func (s *HelloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	var message []string;
	for{
		req, err := stream.Recv();
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessageList{Messages: message})
		}
		if err != nil{
			log.Fatalf("Error Recive Streaming Names :%v", err)
			return err;
		}
		log.Printf("Got The Request Name : %v", req.Name)
		message = append(message, "Hello " + req.Name)
	}

}
