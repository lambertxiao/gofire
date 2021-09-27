package core

import (
	"gofire/iface"
	"log"
)

const DefaultMsgQueueSize = 1024

type FireClient struct {
	server       string
	msgChannel   chan iface.IMsg
	msgQueueSize int

	conn   iface.IConn
	connG  iface.IConnGenerator
	pcodec iface.IPacketCodec
	mcodec iface.IMsgCodec
}

func NewClient(
	server string,
	connG iface.IConnGenerator,
	pcodec iface.IPacketCodec,
	mcodec iface.IMsgCodec,
) iface.IClient {
	c := &FireClient{
		server: server,
		connG:  connG,
		pcodec: pcodec,
		mcodec: mcodec,
	}
	return c
}

func (c *FireClient) SetMsgQueueSize(size int) {
	c.msgQueueSize = size
}

func (c *FireClient) WaitMsg() {
	go func() {
		for {
			data, err := c.pcodec.Decode(c.conn)
			if err != nil {
				log.Println("pcodec decode err: ", err)
				return
			}

			msg, err := c.mcodec.Decode(data)
			if err != nil {
				log.Println("mcodec decode err: ", err)
				return
			}

			c.msgChannel <- msg
		}
	}()
}

func (c *FireClient) Send(msg iface.IMsg) (iface.IMsg, error) {
	conn, err := c.connG.Gen()
	if err != nil {
		return nil, err
	}

	ssm := NewMsgSSM()
	ssm.Add(1)
	log.Println(0)
	stream := NewClientStream(conn, c, ssm)
	stream.Write(msg)
	log.Println(1)
	stream.Flow()
	log.Println("wait done")
	resp := ssm.Done()

	return resp, nil
}

func (c *FireClient) OnMsg() <-chan iface.IMsg {
	return c.msgChannel
}
