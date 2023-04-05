package core

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var defaultTimeout = 6 * time.Second

type DefaultClient struct {
	cg      ConnGenerator
	mq      MsgQueue
	pcodec  PacketCodec
	mcodec  MsgCodec
	timeout time.Duration
	ssmPool *sync.Map
	// 通过client统计超时的消息，在超时达到一定阈值时，会尝试建立新的链接
	maxConnCnt int
	transports []Transport
	inflightm  *sync.Map // store inflight messages
}

func NewClient(
	cg ConnGenerator,
	pcodec PacketCodec,
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
		transports: []Transport{},
		inflightm:  &sync.Map{},
	}

	err := c.buildTransport()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *DefaultClient) buildTransport() error {
	conn, err := c.cg.Gen()
	if err != nil {
		return err
	}

	tp := NewTransport(conn, c.pcodec, c.mcodec)
	tp.SetMsgCB(c.OnMsg)
	tp.Open()
	c.transports = append(c.transports, tp)
	return nil
}

func (c *DefaultClient) OnMsg(tp Transport, msg Msg) {
	val, ok := c.inflightm.LoadAndDelete(msg.GetID())
	if !ok {
		logrus.Info("cannot find inflight msg:", msg.GetID())
		return
	}

	im := val.(*InflightMsg)
	im.Resp = msg
	im.Done()
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

func (c *DefaultClient) selectTransport() (Transport, error) {
	if len(c.transports) == 0 {
		err := c.buildTransport()
		if err != nil {
			return nil, err
		}
	}

	if !c.transports[0].IsActive() {
		c.transports = c.transports[1:]
		err := c.buildTransport()
		if err != nil {
			return nil, err
		}
	}

	return c.transports[0], nil
}

func (c *DefaultClient) SyncSend(msg Msg) (ret Msg, err error) {
	timeout := msg.GetTimeout()
	if timeout == 0 {
		timeout = c.getTimeout()
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	im := newInflightMsg()
	c.inflightm.Store(msg.GetID(), im)

	tp, err := c.selectTransport()
	if err != nil {
		return nil, err
	}

	tp.SendMsg(msg)

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("msg %s send time out", msg.GetID())
	case <-im.Wait():
		return im.Resp, nil
	}
}

func (c *DefaultClient) AsyncSend(msg Msg, cb MsgCallback) error {
	timeout := msg.GetTimeout()
	if timeout == 0 {
		timeout = c.getTimeout()
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	im := newInflightMsg()
	c.inflightm.Store(msg.GetID(), im)
	tp, err := c.selectTransport()
	if err != nil {
		return err
	}

	tp.SendMsg(msg)

	// 由一个统一的去处理超时和回复
	go func() {
		select {
		case <-ctx.Done():
			cb.Timeout(msg)
		case <-im.Wait():
			cb.Success(msg)
		}
	}()

	return nil
}
