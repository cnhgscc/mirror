package cregistry

import (
	"fmt"
	"sync"

	"github.com/hashicorp/consul/api"
)

var (
	crs sync.Map
)

// NewCRegistry new cregistry
func NewCRegistry(scope string, consul ...string) (*CRegistry, error) {
	r, ok := crs.Load(scope)
	if ok {
		return r.(*CRegistry), nil
	}
	if len(consul) == 0 {
		return nil, fmt.Errorf("%s not found", scope)
	}
	client, err := NewClient(consul[0])
	if err != nil {
		return nil, err
	}
	cr := &CRegistry{C: client, N: scope}
	crs.Store(scope, cr)
	return cr, nil
}

type CRegistry struct {
	N  string
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
