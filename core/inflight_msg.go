package core

import (
	"sync"
)

// represent a inflight msg
type InflightMsg struct {
	sync.WaitGroup
	Resp Msg
	Err  error
}

func newInflightMsg() *InflightMsg {
	m := &InflightMsg{}
	m.Add(1)
	return m
}

func (m *InflightMsg) Done() {
	m.WaitGroup.Done()
}

func (m *InflightMsg) WaitDone() (Msg, error) {
	m.WaitGroup.Wait()
	return m.Resp, nil
}
