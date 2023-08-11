package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

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

	addr := fmt.Sprintf("%v:%v", viper.Get("server.host"), viper.Get("server.port"))
	fmt.Println("serve: " + addr)
	lis, _ := net.Listen("tcp", addr)

	g := grpc.NewServer()
	pb.RegisterGreeterServer(g, new(usecase.GrpcGreater))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = g.Serve(lis)
	}()
	defer g.GracefulStop()

	cr, _ := cregistry.NewCRegistry("cr", cregistry.WithGRPCPort(viper.Get("server.port").(int)))
	cr.Register()
	defer cr.UNRegister()
	<-stop
}
