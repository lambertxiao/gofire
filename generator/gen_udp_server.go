package generator

import (
	"gofire/core"
	"gofire/iface"
	"net"
)

type UDPServerConnGenerator struct {
	conn *net.UDPConn
}

func NewUDPServerConnGenerator(endpoint core.Endpoint) (iface.IConnGenerator, error) {
	g := &UDPServerConnGenerator{}
	addr, err := net.ResolveUDPAddr("udp4", endpoint.String())
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return nil, err
	}

	g.conn = conn
	return g, nil
}

func (g *UDPServerConnGenerator) Gen() (iface.IConn, error) {
	return g.conn, nil
}
