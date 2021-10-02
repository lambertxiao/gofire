package core

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var defaultTimeout = 6 * time.Second

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

	transport.Flow()
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

func (c *FireClient) Send(msg IMsg) (ret IMsg, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.getTimeout())
	defer cancel()
	ch := make(chan bool)
	go func() {
		ret, err = c.transport.RoundTrip(msg)
		ch <- true
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("msg id %s send time out", msg.GetID())
	case <-ch:
		return ret, err
	}
}
