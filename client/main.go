package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/AchrafMoreau/gRPC-in-Go/proto"
)

const (
	port = ":9090"
)
func main(){
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()));
	if err != nil{
		log.Fatalf("Error Connecting %v", err);
	}

	defer conn.Close();

	client := pb.NewGreetServiceClient(conn);
	names := &pb.NamesList{
		Names: []string{"Achraf", "kamala", "Yassin", "Anass"},
	}

	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	// callSayHelloClientStream(client, names);
	callHelloBiDirectionStream(client, names)
}
