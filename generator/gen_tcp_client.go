package generator

import (
	"gofire/core"
	"gofire/iface"
	"net"
)

type TCPClientConnGenerator struct {
	endpoint core.Endpoint
}

func NewTCPClientConnGenerator(endpoint core.Endpoint) iface.IConnGenerator {
	g := &TCPClientConnGenerator{
		endpoint: endpoint,
	}
	return g
}

func (g *TCPClientConnGenerator) Gen() (iface.IConn, error) {
	conn, err := net.Dial("tcp4", g.endpoint.String())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
