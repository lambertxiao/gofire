package generator

import (
	"gofire/core"
	"net"
)

type UDPServerConnGenerator struct {
	conn *WrapUDPConn
	ch   chan bool
}

type WrapUDPConn struct {
	*net.UDPConn
	addr *net.UDPAddr
}

func (c *WrapUDPConn) Read(b []byte) (int, error) {
	n, addr, err := c.UDPConn.ReadFromUDP(b)
	if err != nil {
		return n, err
	}

	c.addr = addr
	return n, nil
}

func (c *WrapUDPConn) Write(b []byte) (int, error) {
	return c.UDPConn.WriteToUDP(b, c.addr)
}

func NewUDPServerConnGenerator(endpoint core.Endpoint) (core.ConnGenerator, error) {
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

	g.conn = &WrapUDPConn{UDPConn: conn}

	return g, nil
}

func (g *UDPServerConnGenerator) Gen() (core.Conn, error) {
	g.ch <- true
	return g.conn, nil
}
