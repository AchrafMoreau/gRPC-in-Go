package main

import (
	"context"
	"io"
	"log"

	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Streaming is Started ...")

	stream, err := client.SayHelloServerStream(context.Background(), names);
	if err != nil{
		log.Fatalf("Could not Stream the Names : %v",err)
	}

	for{
		message, err := stream.Recv();
		if err == io.EOF{
			break;
		}

		if err != nil{
			log.Fatalf("Error While Streaming : %v", err);
		}

		log.Println(message)
	}

	log.Printf("Streaming Ended!")

}
