package core

// represent a inflight msg
type InflightMsg struct {
	Resp Msg
	Err  error
	ch   chan struct{}
}

func newInflightMsg() *InflightMsg {
	m := &InflightMsg{
		ch: make(chan struct{}),
	}
	return m
}

func (m *InflightMsg) Done() {
	m.ch <- struct{}{}
}

func (m *InflightMsg) Wait() chan struct{} {
	return m.ch
}
