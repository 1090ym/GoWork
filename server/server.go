package main

import (
	pb "GoWork/proto"
	"context"
	"google.golang.org/grpc"
	"net"
)

var port string = "8888"

type GreeterServer struct {
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloApply, error) {
	return &pb.HelloApply{Message: "Test Server"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)

}
