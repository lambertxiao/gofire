package gofire

import (
	"gofire/iface"
	"io"
	"log"
	"net"
)

type FireConn struct {
	conn net.Conn
}

func NewFireConn(conn net.Conn) iface.IConn {
	c := &FireConn{conn: conn}
	return c
}

func (c *FireConn) Handle() {
	go c.ReadLoop()
	go c.WriteLoop()
}

func (c *FireConn) ReadLoop() {
	for {
		headData := make([]byte, HeaderLength)

		if _, err := io.ReadFull(c.conn, headData); err != nil {
			log.Println("read msg head error", err)
			return
		}

		msg := NewMsg()
		if err := msg.LoadHead(headData); err != nil {
			log.Println("read msg head error", err)
			return
		}

		payloadData := make([]byte, msg.GetPayloadLen())
		if _, err := io.ReadFull(c.conn, payloadData); err != nil {
			log.Println("read msg payload error", err)
			return
		}

		msg.SetPayload(payloadData)
		// go msg.execHandler()
	}
}

func (c *FireConn) WriteLoop() {
}

func (c *FireConn) GetMsg() (iface.IMsg, error) {
	return nil, nil
}
