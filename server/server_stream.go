package main

import (
	pb "grpc_go/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList,stream pb.GreetSevice_SayHelloServerStreamingServer)error{

	log.Printf("got requests with name %v", req.Names)
	for _,name := range req.Names{
		
		res :=&pb.HelloResponse{Message: "Hello"+name}
		if err:=stream.Send(res);err!=nil{
			return err
		}
		time.Sleep(2*time.Second)
	}
	return nil

}


