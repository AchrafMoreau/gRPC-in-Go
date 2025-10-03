package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Println("Client Start Streaming");

	stream, err := client.SayHelloClientStream(context.Background());
	if err != nil{
		log.Println("Could Not Send Names");
	}

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil{
			log.Printf("Error While Sending the name Stream : %v", err);
		}

		time.Sleep(2 * time.Second)
	}

	log.Printf("Names was Send successfuly")
	res, err := stream.CloseAndRecv();
	if err != nil{
		log.Printf("Error While Reciving the name Stream : %v", err);
	}

	log.Printf("%v", res.Messages)

}
