package main

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"mirror/pkg/cmdargs"
	"mirror/pkg/cregistry"
)

func init() {
	cmdargs.Init()
}

func main() {

	Init()
	defer Run()

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
