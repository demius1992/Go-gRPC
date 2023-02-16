package main

import (
	"context"
	//pb "github.com/demius1992/Go-gRPC/proto"
	pb "github.com/demius1992/Go-gRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil

}
