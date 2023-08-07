package cregistry

import (
	"sync"

	"github.com/hashicorp/consul/api"
)

var (
	crs sync.Map
)

// NewCRegistry new cregistry
func NewCRegistry(scope string, consul string) (*CRegistry, error) {
	r, ok := crs.Load(scope)
	if ok {
		return r.(*CRegistry), nil
	}
	client, err := NewClient(consul)
	if err != nil {
		return nil, err
	}
	cr := &CRegistry{C: client}
	crs.Store(scope, cr)
	return cr, nil
}

type CRegistry struct {
	C  *api.Client
	ID string
}

func (cr *CRegistry) Register() {
	go func() {
		_ = register(cr)
	}()

}
func (cr *CRegistry) UNRegister() {
	_ = unregister(cr)
}
