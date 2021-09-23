package gofire

import (
	"fmt"
	"gofire/iface"
	"net"
)

type FireClient struct {
	server  string
	network string
	conn    net.Conn
}

func NewClient(server string) iface.IClient {
	c := &FireClient{
		server:  server,
		network: "tcp4",
	}
	return c
}

func (c *FireClient) Connect() error {
	conn, err := net.Dial(c.network, c.server)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

func (c *FireClient) Send(msg iface.IMsg) error {
	data, err := msg.Pack()
	if err != nil {
		return err
	}

	_, err = c.conn.Write(data)
	if err != nil {
		return err
	}

	fmt.Println(data)
	// c.conn.Read()

	return nil
}
