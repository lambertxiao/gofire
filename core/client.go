package core

import (
	"context"
	"errors"
	"time"
)

type FireClient struct {
	transport ITransport
}

func NewClient(
	transport ITransport,
) IClient {
	c := &FireClient{
		transport: transport,
	}

	return c
}

func (c *FireClient) Send(msg IMsg) (IMsg, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := make(chan bool)
	var ret IMsg
	var err error

	go func() {
		ret, err = c.transport.RoundTrip(msg)
		ch <- true
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("send time out")
	case <-ch:
		return ret, err
	}
}
