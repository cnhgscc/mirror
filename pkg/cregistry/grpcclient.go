package cregistry

import (
	"google.golang.org/grpc"
	"sync"
)

var (
	GRPCConns sync.Map
)

type ClientConn struct {
	*grpc.ClientConn
	ServiceID string
}

func (s *ClientConn) Close() {
	GRPCConns.Delete(s.ServiceID)
	_ = s.ClientConn.Close()
}
