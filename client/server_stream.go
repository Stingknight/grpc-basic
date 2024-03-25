package main

import (
	"context"
	pb "grpc_go/proto"
	"io"
	"log"
)

func callSayHelloServeStream(client pb.GreetSeviceClient,names *pb.NamesList){
	log.Printf("Streaming started")

	stream,err := client.SayHelloServerStreaming(context.Background(),names)
	if err!=nil{
		log.Fatalf("Error: %v",err)
		return
	}

	for {
		message,err := stream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("Error: %v",err)

		}
		log.Println(message)
	}
	log.Printf("Streaming finished")

}