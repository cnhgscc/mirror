package usecase

import (
	"context"

	"github.com/cnhgscc/mirror/pkg/pb"
)

var (
	_ pb.GreeterServer = (*GrpcGreater)(nil)
)

type GrpcGreater struct {
	pb.GreeterServer
}

func (g GrpcGreater) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: request.Name + ":1"}, nil
}

func (g GrpcGreater) SayHelloAgain(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: request.Name + "2"}, nil
}
