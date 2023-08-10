package cregistry

import (
	"net"
)

func ServiceAddr() string {
	var ip = "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipNet.IP.To4() == nil {
			continue
		}
		tmp := ipNet.IP.String()
		if tmp == "127.0.0.1" {
			continue
		}
		ip = tmp
	}
	return ip
}
