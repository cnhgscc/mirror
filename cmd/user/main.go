package main

import (
	"net"

	"google.golang.org/grpc"

	"mirror/pkg/build"
	"mirror/pkg/cmdargs"
	"mirror/pkg/cregistry"
	"mirror/pkg/pb"

	"mirror/internal/pkg/usecase"
)

func init() {
	cmdargs.Init()
}

func main() {

	cr, err := cregistry.NewCRegistry(build.CMDName(), "127.0.0.1:9000")
	if err != nil {
		return
	}
	cr.Register()

	lis, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}

	g := grpc.NewServer()
	pb.RegisterGreeterServer(g, new(usecase.GrpcGreater))
	_ = g.Serve(lis)

}
