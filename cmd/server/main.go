package main

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/cnhgscc/mirror/pkg/cmdargs"
	"github.com/cnhgscc/mirror/pkg/cregistry"
	"github.com/cnhgscc/mirror/pkg/pb"

	"github.com/cnhgscc/mirror/internal/pkg/usecase"
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
