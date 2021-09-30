package generator

import (
	"gofire/core"
	"net"
)

type TCPClientConnGenerator struct {
	endpoint core.Endpoint
}

func NewTCPClientConnGenerator(endpoint core.Endpoint) (core.IChannelGenerator, error) {
	g := &TCPClientConnGenerator{
		endpoint: endpoint,
	}
	return g, nil
}

func (g *TCPClientConnGenerator) Gen() (core.IChannel, error) {
	conn, err := net.Dial("tcp4", g.endpoint.String())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
