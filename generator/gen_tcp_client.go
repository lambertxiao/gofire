package generator

import (
	"gofire/core"
	"net"
)

type TCPClientConnGenerator struct {
	endpoint core.Endpoint
}

func NewTCPClientConnGenerator(endpoint core.Endpoint) core.ConnGenerator {
	g := &TCPClientConnGenerator{
		endpoint: endpoint,
	}
	return g
}

func (g *TCPClientConnGenerator) Gen() (core.Conn, error) {
	conn, err := net.Dial("tcp4", g.endpoint.String())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
