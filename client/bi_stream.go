package main

import (
	"context"
	pb "grpc_go/proto"
	"io"
	"log"
	"time"
)

func callHelloBirectionalStream(client pb.GreetSeviceClient,names *pb.NamesList){
	log.Printf("Birectinal Streaming started")
	
	stream,err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err!=nil{
		log.Fatalf("could not send name %v",err)

	}

	var waitC chan struct{} = make(chan struct{})

	go func(){
		for  {
			message,err := stream.Recv()
			if err!=io.EOF{
				break
			}
			if err!=nil{
				log.Fatalf("error while streaming %v",err)

			}
			log.Println(message)

		}
		close(waitC)
	}()

	for _,name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
				
		}
		if err :=  stream.Send(req);err!=nil{
			log.Fatalf("error while sending %v",err)
		}
		time.Sleep(2* time.Second)
		stream.CloseSend()
		<-waitC

		log.Printf("Bidirectinal streaming is finished")
	}
}