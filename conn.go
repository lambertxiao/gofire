package gofire

import (
	"context"
	"gofire/iface"
	"io"
	"log"
	"net"
)

type FireConn struct {
	conn       net.Conn
	server     iface.IServer
	ctx        context.Context
	cancel     context.CancelFunc
	msgChannel chan []byte
}

func NewFireConn(conn net.Conn, server iface.IServer) iface.IConn {
	ctx, cancel := context.WithCancel(context.Background())
	c := &FireConn{
		conn:       conn,
		server:     server,
		ctx:        ctx,
		cancel:     cancel,
		msgChannel: make(chan []byte),
	}
	return c
}

func (c *FireConn) Handle() {
	go c.ReadLoop()
	go c.WriteLoop()
}

func (c *FireConn) Close() {
	c.conn.Close()
	close(c.msgChannel)
}

func (c *FireConn) ReadLoop() {
	defer c.Close()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			headData := make([]byte, HeaderLength)

			if _, err := io.ReadFull(c.conn, headData); err != nil {
				log.Println("read head data error", err)
				return
			}

			msg := &FireMsg{}
			if err := msg.UnpackHead(headData); err != nil {
				log.Println("read msg head error", err)
				return
			}

			payloadData := make([]byte, msg.GetPayloadLen())
			if _, err := io.ReadFull(c.conn, payloadData); err != nil {
				log.Println("read msg payload error", err)
				return
			}

			msg.SetPayload(payloadData)
			handler := c.server.GetActionHandler(msg.GetAction())
			if handler == nil {
				log.Println("not support action")
				c.WriteMsg(NewMsg(0, 0, []byte("not support action")))
				return
			}

			req := iface.Request{
				Conn: c,
				Msg:  msg,
			}

			log.Println("server receive msg from conn", msg.ID)
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
			_, err := c.conn.Write(msgData)
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
