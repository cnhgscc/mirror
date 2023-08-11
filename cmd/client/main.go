package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cnhgscc/mirror/pkg/cmdargs"
	"github.com/cnhgscc/mirror/pkg/cregistry"
	"github.com/cnhgscc/mirror/pkg/pb"
)

func init() {
	cmdargs.Init()
}

func Run() {

	go func() {
		for true {
			cr, _ := cregistry.NewCRegistry("cr")
			gs, err := cr.GS("server")
			if err != nil {
				return
			}

			client := pb.NewGreeterClient(gs)
			args := &pb.HelloRequest{Name: "3123"}
			reply, _ := client.SayHello(context.Background(), args)
			fmt.Println(reply)
			time.Sleep(3 * time.Second)
		}
	}()
}

func main() {

	addr := fmt.Sprintf("%v:%v", viper.Get("server.host"), viper.Get("server.port"))
	fmt.Println("serve: " + addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGALRM)
	g := grpc.NewServer()
	go func() {
		_ = g.Serve(lis)
	}()
	defer g.GracefulStop()

	cr, err := cregistry.NewCRegistry("cr", cregistry.WithGRPCPort(viper.Get("server.port").(int)))
	if err != nil {
		return
	}
	cr.Register()
	defer cr.UNRegister()

	Run()

	<-stop
}
