package generator

import (
	"gofire/core"
	"gofire/iface"
	"net"
)

type TCPServerConnGenerator struct {
	listener *net.TCPListener
}

func NewTCPServerConnGenerator(endpoint core.Endpoint) (iface.IConnGenerator, error) {
	g := &TCPServerConnGenerator{}
	addr, err := net.ResolveTCPAddr("tcp4", endpoint.String())
	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		return nil, err
	}

	g.listener = listener
	return g, nil
}

func (g *TCPServerConnGenerator) Gen() (iface.IConn, error) {
	for {
		conn, err := g.listener.AcceptTCP()
		if err != nil {
			return nil, err
		}

		return conn, nil
	}
}
