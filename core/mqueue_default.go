package core

type DefaultMsgQueue struct {
	ch chan Msg
}

func NewDefaultMsgQueue(cap uint32) MsgQueue {
	q := &DefaultMsgQueue{
		ch: make(chan Msg, cap),
	}
	return q
}

func (q *DefaultMsgQueue) Push(msg Msg) {
	q.ch <- msg
}

func (q *DefaultMsgQueue) Pop() Msg {
	return <-q.ch
}
