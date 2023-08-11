package cregistry

import (
	"github.com/hashicorp/consul/api"
)

func NewClient(addr string) (*api.Client, error) {
	c := api.DefaultConfig()
	c.Address = addr

	client, err := api.NewClient(c)
	if err != nil {
		return nil, err
	}
	return client, nil
}
