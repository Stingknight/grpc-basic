package main

import (
	"context"
	pb "grpc_go/proto"
)


func (s *helloServer) SayHello(ctx context.Context,req *pb.NoParam)(*pb.HelloResponse,error){
	return &pb.HelloResponse{
		Message: "hello"},
		nil
}


