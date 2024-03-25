package main

import (
	pb "grpc_go/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetSevice_SayHelloBiDirectionalStreamingServer) error {
	for {
		req,err := stream.Recv()
		if err!=io.EOF{
			return err
		}
		if err!=nil{
			return err
		}

		log.Printf("Got a request with name: %v",req.Name)
		res := &pb.HelloResponse{
			Message: "Hello, "+req.Name,
		}
		if err := stream.Send(res);err!=nil{
			return err	
		}	
	}

}
