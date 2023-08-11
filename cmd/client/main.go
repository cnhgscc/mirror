package main

import (
	"context"
	"fmt"
	"github.com/cnhgscc/mirror/pkg/pb"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/cnhgscc/mirror/pkg/cmdargs"
	"github.com/cnhgscc/mirror/pkg/cregistry"
)

func init() {
	cmdargs.Init()
}

func main() {

	Init()
	defer Run()

	cr, _ := cregistry.NewCRegistry("cr")
	gs, _ := cr.GS("grpc")

	client := pb.NewGreeterClient(gs)
	args := &pb.HelloRequest{Name: "3123"}
	reply, _ := client.SayHello(context.Background(), args)
	fmt.Println(reply)
}

func Init() {

	cr, err := cregistry.NewCRegistry("cr", cregistry.WithHTTPPort(viper.Get("server.port").(int)))
	if err != nil {
		return
	}
	cr.Register()
	defer cr.UNRegister()

}

func Run() {

	addr := fmt.Sprintf("%v:%v", viper.Get("server.host"), viper.Get("server.port"))
	fmt.Println("serve: " + addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	g := grpc.NewServer()
	_ = g.Serve(lis)

}
