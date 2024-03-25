package main

import (
	
	pb "grpc_go/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port =":9090"

)

type helloServer struct{
	pb.GreetSeviceServer
}

func main(){

	lis,err := net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("error: %v",err)
		return
	}
	grpcServer := grpc.NewServer()

	pb.RegisterGreetSeviceServer(grpcServer,&helloServer{})

	log.Printf("Server is listening on port %v",lis.Addr())

	if err := grpcServer.Serve(lis);err!=nil{
		log.Fatalf("Error: %v",err)
		return
	}

	

}