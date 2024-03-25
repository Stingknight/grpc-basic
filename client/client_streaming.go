package main

import (
	"context"
	"log"
	"time"
	pb "grpc_go/proto"
)

func callSayHelloClientStream(client pb.GreetSeviceClient,names *pb.NamesList){

	log.Println("Client streaming started")

	stream,err := client.SayHelloClientStreaming(context.Background())
	if err!=nil{
		log.Fatalf("Error :%v",err)
		return
	}
	for _,name := range names.Names{

		req := &pb.HelloRequest{Name:name,}

		if err := stream.Send(req);err!=nil{
			log.Fatalf("Error :%v\n",err)
			return
		}
		log.Printf("send the name: %s",name)
		time.Sleep(2*time.Second)
	}
	res,err := stream.CloseAndRecv()
	log.Println("Client streaming finished")
	
	if err!=nil{
		log.Fatalf("Error :%v\n",err)
		return
	}
	log.Printf("%v\n",res.Messages)
}