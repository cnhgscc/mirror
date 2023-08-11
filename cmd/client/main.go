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
	services := cr.Services("grpc")
	for _, srv := range services {
		host := srv.ServiceAddress
		port, ok := srv.ServiceMeta[cregistry.GRPCPort]
		if !ok || port == "" {
			continue
		}

		dial, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
		if err != nil {
			return
		}
		client := pb.NewGreeterClient(dial)
		args := &pb.HelloRequest{Name: "3123"}
		reply, err := client.SayHello(context.Background(), args)
		if err != nil {
			return
		}
		fmt.Println(reply)

	}

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
