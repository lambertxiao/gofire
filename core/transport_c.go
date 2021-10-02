package core

import (
	"log"
	"sync"
)

type Transport struct {
	ch      IChannel
	pcodec  IPacketCodec
	mcodec  IMsgCodec
	mq      IMsgQueue
	ssmPool *sync.Map
}

func NewTransport(ch IChannel, pcodec IPacketCodec, mcodec IMsgCodec) ITransport {
	t := &Transport{
		ch:      ch,
		pcodec:  pcodec,
		mcodec:  mcodec,
		mq:      NewMsgQueue(128),
		ssmPool: &sync.Map{},
	}
	return t
}

func (t *Transport) Flow() {
	go t.WriteLoop()
	go t.ReadLoop()
}

func (t *Transport) WriteLoop() {
	for {
		msg := t.mq.Pop()
		payload, err := t.mcodec.Encode(msg)
		if err != nil {
			log.Println(err)
			continue
		}

		if err = t.pcodec.Encode(payload, t.ch); err != nil {
			log.Println(err)
			continue
		}
	}
}

func (t *Transport) ReadLoop() {
	for {
		repayload, err := t.pcodec.Decode(t.ch)
		if err != nil {
			log.Println(err)
			continue
		}

		msg, err := t.mcodec.Decode(repayload)
		if err != nil {
			log.Println(err)
			continue
		}

		_ssm, ok := t.ssmPool.LoadAndDelete(msg.GetID())
		if !ok {
			log.Println("cannot find ssm for msg:", msg.GetID())
			continue
		}

		ssm := _ssm.(*MsgSSM)
		ssm.Resp = msg
		ssm.Done()
	}
}

func (t *Transport) RoundTrip(msg IMsg) (IMsg, error) {
	// msg bind ssm
	ssm := NewMsgSSM()
	ssm.Go()
	t.ssmPool.Store(msg.GetID(), ssm)
	t.mq.Push(msg)
	return ssm.Return()
}
