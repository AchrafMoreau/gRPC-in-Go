package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second);
	defer cancel();

	res, err := client.SayHello(ctx, &pb.NoParam{});
	if err != nil{
		log.Fatalf("Could Not Greet %v", err)
	}

	log.Printf("This's the Answaer : %v", res.Message)
}
