package gofire

import (
	"gofire/iface"
	"log"
	"net"
)

type TCPServerConnGenerator struct {
	listener *net.TCPListener
}

func NewTCPServerConnGenerator(endpoint Endpoint) (iface.IConnGenerator, error) {
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

func (g *TCPServerConnGenerator) Gen() iface.IConn {
	for {
		conn, err := g.listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}

		return conn
	}
}
