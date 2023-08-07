package cregistry

import (
	"github.com/hashicorp/consul/api"
)

func NewClient(addr string) (*api.Client, error) {
	c := api.DefaultConfig()
	c.Address = "127.0.0.1:8500"

	client, err := api.NewClient(c)
	if err != nil {
		return nil, err
	}
	return client, nil
}
