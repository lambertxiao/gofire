package core

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var defaultTimeout = 6 * time.Second

type DefaultClient struct {
	cg      ConnGenerator
	mq      MsgQueue
	pcodec  IPacketCodec
	mcodec  MsgCodec
	timeout time.Duration
	ssmPool *sync.Map
	// 通过client统计超时的消息，在超时达到一定阈值时，会尝试建立新的链接
	maxConnCnt int
	transports []ClientTransport
}

func NewClient(
	cg ConnGenerator,
	pcodec IPacketCodec,
	mcodec MsgCodec,
	mq MsgQueue,
	maxConnCnt int,
) (Client, error) {
	c := &DefaultClient{
		cg:         cg,
		mq:         mq,
		pcodec:     pcodec,
		mcodec:     mcodec,
		ssmPool:    &sync.Map{},
		maxConnCnt: maxConnCnt,
		transports: []ClientTransport{},
	}

	conn, err := cg.Gen()
	if err != nil {
		return nil, err
	}

	tp := NewClientTransport(conn, c.pcodec, c.mcodec)
	c.transports = append(c.transports, tp)
	return c, nil
}

func (c *DefaultClient) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *DefaultClient) getTimeout() time.Duration {
	if c.timeout == 0 {
		return defaultTimeout
	}

	return c.timeout
}

func (c *DefaultClient) selectTransport() ClientTransport {
	return c.transports[0]
}

func (c *DefaultClient) SyncSend(msg Msg) (ret Msg, err error) {
	timeout := msg.GetTimeout()
	if timeout == 0 {
		timeout = c.getTimeout()
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan bool)
	go func() {
		ret, err = c.selectTransport().RoundTrip(msg)
		ch <- true
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("msg id %s send time out", msg.GetID())
	case <-ch:
		return ret, err
	}
}

func (c *DefaultClient) AsyncSend(msg Msg, cb MsgCB) error {
	go func() {
		ret, err := c.selectTransport().RoundTrip(msg)
		cb(ret, err)
	}()

	return nil
}
