package main

import (
	"log"
	 pb "grpc_go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


const (
	port = ":9090"
)

func main(){
	conn,err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewBundle().TransportCredentials()))
	if err!=nil{
		log.Fatalf("Error: %v", err)
		return
	}

	defer conn.Close()

	client := pb.NewGreetSeviceClient(conn)

	names := &pb.NamesList{Names:[]string{"john","jane","robert"}}

	// callSayHello(client)
	
	// callSayHelloServeStream(client,names)

	// callSayHelloClientStream(client,names)
	callHelloBirectionalStream(client ,names)

}