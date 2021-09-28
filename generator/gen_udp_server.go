package generator

import (
	"fmt"
	"gofire/core"
	"gofire/iface"
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
	fmt.Println(c.addr)
	return c.UDPConn.WriteToUDP(b, c.addr)
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

	g.conn = &WrapUDPConn{UDPConn: conn}

	return g, nil
}

func (g *UDPServerConnGenerator) Gen() (iface.IConn, error) {
	g.ch <- true
	return g.conn, nil
}
