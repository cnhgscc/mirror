package main

import (
	"net"

	"google.golang.org/grpc"

	"mirror/pkg/pb"

	"mirror/internal/pkg/usecase"
)

func main() {

	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		return
	}

	g := grpc.NewServer()
	pb.RegisterGreeterServer(g, new(usecase.GrpcGreater))
	_ = g.Serve(lis)

}
