package gofire

import (
	"context"
	"gofire/iface"
	"io"
	"log"
	"net"
)

type FireConn struct {
	net.Conn
	server     iface.IServer
	ctx        context.Context
	cancel     context.CancelFunc
	msgChannel chan []byte
}

func NewFireConn(conn net.Conn, server iface.IServer) iface.IConn {
	ctx, cancel := context.WithCancel(context.Background())
	c := &FireConn{
		Conn:   conn,
		server: server,
		ctx:    ctx,
		cancel: cancel,
	}
	return c
}

func (c *FireConn) Handle() {
	go c.ReadLoop()
	go c.WriteLoop()
}

func (c *FireConn) ReadLoop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			headData := make([]byte, HeaderLength)

			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				log.Println("read msg head error", err)
				return
			}

			msg := NewMsg()
			if err := msg.UnpackHead(headData); err != nil {
				log.Println("read msg head error", err)
				return
			}

			payloadData := make([]byte, msg.GetPayloadLen())
			if _, err := io.ReadFull(c.Conn, payloadData); err != nil {
				log.Println("read msg payload error", err)
				return
			}

			msg.SetPayload(payloadData)

			handler := c.server.GetActionHandler(msg.GetAction())
			if handler == nil {
				log.Println("not support action")
				return
			}

			req := iface.Request{
				Conn: c,
				Msg:  msg,
			}
			go handler.Do(req)
		}
	}
}

func (c *FireConn) WriteLoop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case msgData := <-c.msgChannel:
			_, err := c.Write(msgData)
			if err != nil {
				log.Println("write msg data to connection error", err)
				return
			}
		}
	}
}

func (c *FireConn) WriteMsg(msg iface.IMsg) {
	msgData, err := msg.Pack()
	if err != nil {
		log.Println("pack msg error", err)
		return
	}

	c.msgChannel <- msgData
}
