package cregistry

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func register(cr *CRegistry) error {

	osip := ServiceAddr()
	id := fmt.Sprintf("%s@%s", cr.Name, osip)

	err := cr.C.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      id,
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
	osip := ServiceAddr()
	id := fmt.Sprintf("%s@%s", cr.Name, osip)
	return cr.C.Agent().ServiceDeregister(id)
}
