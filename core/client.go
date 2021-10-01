package core

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

var defaultTimeout = 2 * time.Second

type FireClient struct {
	transport ITransport
	mq        IMsgQueue
	timeout   time.Duration
	ssmPool   *sync.Map
}

func NewClient(
	transport ITransport,
) IClient {
	c := &FireClient{
		transport: transport,
		mq:        NewMsgQueue(128),
		ssmPool:   &sync.Map{},
	}

	go c.loop()
	return c
}

func (c *FireClient) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *FireClient) getTimeout() time.Duration {
	if c.timeout == 0 {
		return defaultTimeout
	}

	return c.timeout
}

func (c *FireClient) loop() {
	for {
		msg := c.mq.Pop()
		ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout())
		ch := make(chan bool)

		_ssm, ok := c.ssmPool.Load(msg.GetID())
		if !ok {
			log.Println("cannot find ssm for msg:", msg.GetID())
			continue
		}
		ssm := _ssm.(*MsgSSM)

		go func() {
			ret, err := c.transport.RoundTrip(msg)
			if err != nil {
				log.Println("transport roundtrip msg failed")
				return
			}

			ssm.Err = err
			ssm.Resp = ret
			ssm.Done()
		}()

		select {
		case <-ctx.Done():
			ssm.Err = errors.New("send time out")
		case <-ch:
			cancel()
		}
	}
}

func (c *FireClient) Send(msg IMsg) (IMsg, error) {
	ssm := NewMsgSSM()
	c.mq.Push(msg)
	ssm.Go()
	c.ssmPool.Store(msg.GetID(), ssm)
	return ssm.Return()
}
