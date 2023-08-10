package cregistry

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// 1. 使用 viper 获取服务的端口
// 2. 动态获取服务的ip
func register(cr *CRegistry) error {

	osip := "127.0.0.1"
	cr.ID = fmt.Sprintf("%s@%s", cr.Name, osip)

	err := cr.C.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      cr.ID,
		Name:    cr.Name,
		Address: osip,
		Port:    cr.HTTPPort,
		Meta:    cr.Meta,
		Check: &api.AgentServiceCheck{
			CheckID:  "tcp",
			TCP:      fmt.Sprintf("%s:%d", osip, cr.HTTPPort),
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
