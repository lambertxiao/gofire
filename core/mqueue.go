package core

type MsgQueue struct {
	ch chan IMsg
}

func NewMsgQueue(cap uint32) IMsgQueue {
	q := &MsgQueue{
		ch: make(chan IMsg, cap),
	}
	return q
}

func (q *MsgQueue) Push(msg IMsg) {
	q.ch <- msg
}

func (q *MsgQueue) Pop() IMsg {
	return <-q.ch
}
