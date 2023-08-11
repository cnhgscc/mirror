package usecase

import (
	"context"
	"time"

	"github.com/cnhgscc/mirror/pkg/pb"
)

var (
	_ pb.GreeterServer = (*GrpcGreater)(nil)
)

type GrpcGreater struct {
	pb.GreeterServer
}

func (g GrpcGreater) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: request.Name + ":" + time.Now().Format(time.RFC3339)}, nil
}

func (g GrpcGreater) SayHelloAgain(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: request.Name + "2"}, nil
}
