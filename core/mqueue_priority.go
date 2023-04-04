package core

type PriorityMsgQueue struct {
	ch chan Msg
}

func NewPriorityMsgQueue(cap uint32) MsgQueue {
	q := &DefaultMsgQueue{
		ch: make(chan Msg, cap),
	}
	return q
}

func (q *PriorityMsgQueue) Push(msg Msg) {
	q.ch <- msg
}

func (q *PriorityMsgQueue) Pop() Msg {
	return <-q.ch
}
