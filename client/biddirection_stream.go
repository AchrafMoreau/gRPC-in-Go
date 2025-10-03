package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

func callHelloBiDirectionStream(client pb.GreetServiceClient, names *pb.NamesList) {
	stream, err := client.SayHelloBiDirectionStreaming(context.Background())
	if err != nil {
		log.Println("Could Not Send Names")
	}

	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error Receving the Names :%v", err)
			}

			log.Printf("Got The Name : %v", message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Printf("Error Streaming the Names :%v", err)
		}
		time.Sleep(2 * time.Second)
	}


	stream.CloseSend()
	<- waitc
	log.Println("Streaming Finished !----------")

}
