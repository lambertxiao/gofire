package gofire

import (
	"gofire/iface"
	"net"
)

type TCPClientConnGenerator struct {
	address string
}

func (g TCPClientConnGenerator) Gen() (iface.IConn, error) {
	conn, err := net.Dial("tcp4", g.address)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
