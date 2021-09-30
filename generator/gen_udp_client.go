package generator

import (
	"gofire/core"
	"net"
)

type UDPClientConnGenerator struct {
	addr *net.UDPAddr
}

func NewUDPClientConnGenerator(endpoint core.Endpoint) (core.IChannelGenerator, error) {
	g := &UDPClientConnGenerator{}
	addr, err := net.ResolveUDPAddr("udp4", endpoint.String())
	if err != nil {
		return nil, err
	}

	g.addr = addr
	return g, nil
}

func (g *UDPClientConnGenerator) Gen() (core.IChannel, error) {
	conn, err := net.DialUDP("udp4", nil, g.addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
