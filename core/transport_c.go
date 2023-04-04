package core

import (
	"log"
	"sync"
)

type DefaultClientTransport struct {
	conn      Conn
	pcodec    IPacketCodec
	mcodec    MsgCodec
	mq        MsgQueue
	inflightm *sync.Map // store inflight messages
	timeoutm  *sync.Map // store timeout messages
}

func NewClientTransport(conn Conn, pcodec IPacketCodec, mcodec MsgCodec) ClientTransport {
	t := &DefaultClientTransport{
		conn:      conn,
		pcodec:    pcodec,
		mcodec:    mcodec,
		mq:        NewDefaultMsgQueue(128),
		inflightm: &sync.Map{},
		timeoutm:  &sync.Map{},
	}
	t.Open()
	return t
}

func (t *DefaultClientTransport) Open() {
	go t.WriteLoop()
	go t.ReadLoop()
}

func (t *DefaultClientTransport) WriteLoop() {
	for {
		msg := t.mq.Pop()
		payload, err := t.mcodec.Encode(msg)
		if err != nil {
			log.Println(err)
			continue
		}

		if err = t.pcodec.Encode(payload, t.conn); err != nil {
			log.Println(err)
			continue
		}
	}
}

func (t *DefaultClientTransport) ReadLoop() {
	for {
		repayload, err := t.pcodec.Decode(t.conn)
		if err != nil {
			log.Println(err)
			continue
		}

		msg, err := t.mcodec.Decode(repayload)
		if err != nil {
			log.Println(err)
			continue
		}

		_im, ok := t.inflightm.LoadAndDelete(msg.GetID())
		if !ok {
			log.Println("cannot find ssm for msg:", msg.GetID())
			continue
		}

		im := _im.(*InflightMsg)
		im.Resp = msg
		im.Done()
	}
}

func (t *DefaultClientTransport) RoundTrip(msg Msg) (Msg, error) {
	im := newInflightMsg()
	t.inflightm.Store(msg.GetID(), im)
	t.mq.Push(msg)
	return im.WaitDone()
}
