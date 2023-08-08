package cregistry

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"

	"mirror/pkg/build"
)

// 1. 使用 viper 获取服务的端口
// 2. 动态获取服务的ip
func register(cr *CRegistry) error {

	osip := "127.0.0.1"
	port := viper.GetInt("server.port")
	cr.ID = fmt.Sprintf("%s:%d", osip, port)

	err := cr.C.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      cr.ID,
		Name:    build.CMDName(),
		Address: osip,
		Port:    port,
		Check: &api.AgentServiceCheck{
			CheckID:  "g",
			TCP:      fmt.Sprintf("%s:%d", osip, port),
			Timeout:  "1s",
			Interval: "3s",
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func unregister(cr *CRegistry) error {
	return cr.C.Agent().ServiceDeregister(cr.ID)

}
