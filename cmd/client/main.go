package main

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"mirror/pkg/build"
	"mirror/pkg/cmdargs"
	"mirror/pkg/cregistry"
)

func init() {
	cmdargs.Init()
}

func main() {

	Init()
	defer Run()

	cr, _ := cregistry.NewCRegistry(build.CMDName())
	s, _ := cr.S("127.0.0.1:9002")
	fmt.Println(s.Address, s.ID, s.Namespace)

}

func Init() {

	cr, err := cregistry.NewCRegistry(build.CMDName(), "127.0.0.1:9000")
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
