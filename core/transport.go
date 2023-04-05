package core

import (
	"io"

	"github.com/sirupsen/logrus"
)

type DefaultTransport struct {
	conn         Conn
	pcodec       PacketCodec
	mcodec       MsgCodec
	sendq        MsgQueue
	msgCB        MsgRecvCallback
	disconnected bool
}

func NewTransport(conn Conn, pcodec PacketCodec, mcodec MsgCodec) Transport {
	t := &DefaultTransport{
		conn:   conn,
		pcodec: pcodec,
		mcodec: mcodec,
		sendq:  NewDefaultMsgQueue(128),
	}
	return t
}

func (t *DefaultTransport) Open() {
	go t.WriteLoop()
	go t.ReadLoop()
}

func (t *DefaultTransport) WriteLoop() {
	for {
		msg := t.sendq.Pop()
		payload, err := t.mcodec.Encode(msg)
		if err != nil {
			logrus.Info(err)
			continue
		}

		if err = t.pcodec.Encode(payload, t.conn); err != nil {
			if err == io.EOF {
				logrus.Info("transport encode with eof, exit")
				t.disconnected = true
				break
			}
			logrus.Info(err)
			continue
		}
	}
}

func (t *DefaultTransport) ReadLoop() {
	for {
		repayload, err := t.pcodec.Decode(t.conn)
		if err != nil {
			if err == io.EOF {
				logrus.Info("transport decode with eof, exit")
				t.disconnected = true
				break
			}
			logrus.Info(err)
			continue
		}

		msg, err := t.mcodec.Decode(repayload)
		if err != nil {
			logrus.Info(err)
			continue
		}

		if t.msgCB != nil {
			go t.msgCB(t, msg)
		}
	}
}

func (t *DefaultTransport) SendMsg(msg Msg) error {
	t.sendq.Push(msg)
	return nil
}

func (t *DefaultTransport) SetMsgCB(cb MsgRecvCallback) {
	t.msgCB = cb
}

func (t *DefaultTransport) IsActive() bool {
	return !t.disconnected
}
