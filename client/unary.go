package main

import (
	"context"
	pb "grpc_go/proto"
	"log"
	"time"
)


func callSayHello(client pb.GreetSeviceClient){
	ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)

	defer cancel()

	res,err := client.SayHello(ctx,&pb.NoParam{})
	if err!=nil{
		log.Fatalf("Error: %v",err)
		return
	}
	 
	log.Printf("%s",res.Message)
}
