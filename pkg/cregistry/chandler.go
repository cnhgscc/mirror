package cregistry

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/cnhgscc/mirror/pkg/build"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
	N string
	sync.Mutex
	C *api.Client

	Node // 节点信息
}

func (cr *CRegistry) Register() {
	go func() {
		_ = register(cr)
	}()

}
func (cr *CRegistry) UNRegister() {
	_ = unregister(cr)
}

var (
	GRPCNodes sync.Map // GRPC节点信息
)

func (cr *CRegistry) GS(name string) (*ClientConn, error) {
	node, ok := GRPCNodes.Load(name)
	if !ok {
		timer := time.NewTicker(3 * time.Second)
		go func() {
			for range timer.C {
				ss, _, _ := cr.C.Catalog().Service(name, "", nil)
				if len(ss) == 0 {
					return
				}
				GRPCNodes.Store(name, ss)
			}
		}()
		first, _, _ := cr.C.Catalog().Service(name, "", nil)
		if len(first) == 0 {
			GRPCNodes.Store(name, []*api.CatalogService{})
			return nil, fmt.Errorf("%s not found", name)
		}
		GRPCNodes.Store(name, first)
		node, _ = GRPCNodes.Load(name)
	}

	services, ok := node.([]*api.CatalogService)
	if !ok || len(services) == 0 {
		return nil, fmt.Errorf("%s not found", name)
	}

	index := rand.Intn(len(services))
	gs := services[index]
	port, ok := gs.ServiceMeta[GRPCPort]
	if !ok || port == "" {
		return nil, fmt.Errorf("%s not support", name)
	}
	host := gs.ServiceAddress

	cc, ok := GRPCConns.Load(gs.ServiceID)
	if ok {
		return cc.(*ClientConn), nil
	}

	tmp, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	conn := &ClientConn{tmp, gs.ServiceID}
	GRPCConns.Store(gs.ServiceID, conn)
	return conn, nil
}
