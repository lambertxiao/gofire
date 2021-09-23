package gofire

import (
	"gofire/iface"
	"io"
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

func (c *FireClient) SyncSend(msg iface.IMsg) ([]byte, error) {
	data, err := msg.Pack()
	if err != nil {
		return nil, err
	}

	_, err = c.conn.Write(data)
	if err != nil {
		return nil, err
	}

	headData := make([]byte, HeaderLength)
	_, err = io.ReadFull(c.conn, headData)
	if err != nil {
		return nil, err
	}

	respMsg := &FireMsg{}
	respMsg.UnpackHead(headData)
	payload := make([]byte, respMsg.PayloadLength)
	_, err = io.ReadFull(c.conn, payload)

	if err != nil {
		return nil, err
	}

	return payload, nil
}
