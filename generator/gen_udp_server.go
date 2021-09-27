package generator

import (
	"fmt"
	"gofire/core"
	"gofire/iface"
	"net"
)

type UDPServerConnGenerator struct {
	conn *net.UDPConn
	ch   chan bool
}

func NewUDPServerConnGenerator(endpoint core.Endpoint) (iface.IConnGenerator, error) {
	g := &UDPServerConnGenerator{
		ch: make(chan bool, 1),
	}
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
	g.ch <- true
	fmt.Println("kkk")
	return g.conn, nil
}
