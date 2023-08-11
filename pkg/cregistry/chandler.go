package cregistry

import (
	"fmt"
	"sync"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"

	"github.com/cnhgscc/mirror/pkg/build"
)

var (
	crs sync.Map
)

const (
	GRPCPort = "grpc_port"
	HTTPPort = "http_port"
)

// NewCRegistry new cregistry
func NewCRegistry(scope string, opt ...Option) (*CRegistry, error) {

	r, ok := crs.Load(scope)
	if ok {
		return r.(*CRegistry), nil
	}

	node := &Node{
		Registry: "127.0.0.1:8500",

		Name:     build.CMDName(),
		GRPCPort: 7001,
		HTTPPort: viper.GetInt("server.port"),
		Meta: map[string]string{
			GRPCPort: "",
			HTTPPort: "",
		},
	}

	for _, option := range opt {
		option(node)
	}

	client, err := NewClient(node.Registry)
	if err != nil {
		return nil, err
	}

	cr := &CRegistry{C: client, N: scope}
	cr.Node = *node
	crs.Store(scope, cr)
	return cr, nil
}

type CRegistry struct {
	Node

	C *api.Client

	ID string
	N  string
}

func (cr *CRegistry) Register() {
	go func() {
		_ = register(cr)
	}()

}
func (cr *CRegistry) UNRegister() {
	_ = unregister(cr)
}

func (cr *CRegistry) Services(name string) []*api.CatalogService {
	service, meta, err := cr.C.Catalog().Service(name, "", nil)
	fmt.Println(meta, err)
	return service

}
