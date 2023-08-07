package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"mirror/pkg/cregistry"
	"mirror/pkg/pb"

	"mirror/internal/pkg/usecase"
)

func main() {

	go func() {
		err := cregistry.Register()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	lis, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}

	g := grpc.NewServer()
	pb.RegisterGreeterServer(g, new(usecase.GrpcGreater))
	_ = g.Serve(lis)

}
