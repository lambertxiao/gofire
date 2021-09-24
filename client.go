package gofire

import (
	"gofire/iface"
	"io"
	"log"
	"net"
)

const DefaultMsgQueueSize = 1024

type FireClient struct {
	server       string
	network      string
	conn         net.Conn
	msgChannel   chan iface.IMsg
	msgQueueSize int
}

func NewClient(server string) iface.IClient {
	c := &FireClient{
		server:  server,
		network: "tcp4",
	}
	return c
}

func (c *FireClient) Connect() error {
	if c.msgQueueSize != 0 {
		c.msgChannel = make(chan iface.IMsg, c.msgQueueSize)
	} else {
		c.msgChannel = make(chan iface.IMsg, DefaultMsgQueueSize)
	}

	conn, err := net.Dial(c.network, c.server)
	if err != nil {
		return err
	}

	c.conn = conn
	c.WaitMsg()

	return nil
}

func (c *FireClient) SetMsgQueueSize(size int) {
	c.msgQueueSize = size
}

func (c *FireClient) WaitMsg() {
	go func() {
		for {
			headData := make([]byte, HeaderLength)
			_, err := io.ReadFull(c.conn, headData)
			if err != nil {
				log.Println("read head data err: ", err)
				return
			}

			respMsg := &FireMsg{}
			respMsg.UnpackHead(headData)
			payload := make([]byte, respMsg.PayloadLength)
			_, err = io.ReadFull(c.conn, payload)

			if err != nil {
				log.Println("read payload data err: ", err)
				return
			}

			respMsg.SetPayload(payload)
			c.msgChannel <- respMsg
		}
	}()
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

	return nil
}

func (c *FireClient) OnMsg() <-chan iface.IMsg {
	return c.msgChannel
}
