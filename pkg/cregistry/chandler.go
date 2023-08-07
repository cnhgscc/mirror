package cregistry

import (
	"github.com/hashicorp/consul/api"
)

func Register() error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
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
