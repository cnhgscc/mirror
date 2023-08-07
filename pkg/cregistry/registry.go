package cregistry

import (
	"github.com/hashicorp/consul/api"
)

// 1. 使用 viper 获取服务的端口
// 2. 动态获取服务的ip
func register(client *api.Client) error {

	err := client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      "g",
		Name:    "g",
		Address: "127.0.0.1",
		Port:    9000,
		Check: &api.AgentServiceCheck{
			CheckID:  "g",
			TCP:      "127.0.0.1:9000",
			Timeout:  "1s",
			Interval: "3s",
		},
	})
	if err != nil {
		return err
	}
	return nil
}
