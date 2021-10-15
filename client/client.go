package main

import (
	pb "GoWork/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

var port string = "8888"

func SayHello(client pb.GreeterClient) error {
	rep, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "client"})
	log.Printf("client.rep: %s", rep.Message)
	return nil
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
}
