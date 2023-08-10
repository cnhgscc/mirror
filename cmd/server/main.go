package main

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"mirror/pkg/cmdargs"
	"mirror/pkg/cregistry"
	"mirror/pkg/pb"

	"mirror/internal/pkg/usecase"
)

func init() {
	cmdargs.Init()
}

func main() {

	cr, err := cregistry.NewCRegistry("cr", cregistry.WithGRPCPort(viper.Get("server.port").(int)))
	if err != nil {
		return
	}
	cr.Register()
	defer cr.UNRegister()

	addr := fmt.Sprintf("%v:%v", viper.Get("server.host"), viper.Get("server.port"))
	fmt.Println("serve: " + addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	g := grpc.NewServer()
	pb.RegisterGreeterServer(g, new(usecase.GrpcGreater))
	_ = g.Serve(lis)

}
