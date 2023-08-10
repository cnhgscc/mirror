package cregistry

import (
	"fmt"
)

type Option func(node *Node)

type Node struct {
	Registry string

	Name     string
	GRPCPort int
	HTTPPort int
	Meta     map[string]string
}

// WithRegistry consul registry add
func WithRegistry(consul string) Option {
	return func(node *Node) {
		node.Registry = consul
	}
}

// WithName node name default BuildName
func WithName(name string) Option {
	return func(node *Node) {
		node.Name = name
	}
}

// WithGRPCPort grpc port
func WithGRPCPort(port int) Option {
	return func(node *Node) {
		node.GRPCPort = port
		node.Meta["grpc_port"] = fmt.Sprintf("%d", port)
	}
}

// WithHTTPPort http port
func WithHTTPPort(port int) Option {
	return func(node *Node) {
		node.HTTPPort = port
		node.Meta["http_port"] = fmt.Sprintf("%d", port)
	}
}
